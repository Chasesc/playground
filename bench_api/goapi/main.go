package main

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MassGrams   int    `json:"mass_grams"`
}

var Db *sql.DB

func main() {
	connectDb()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})
	r.GET("/items/", getItems)
	r.GET("/items/:item_id", getItemByID)
	r.POST("/items/", postItem)
	r.Run()
}

func getItems(c *gin.Context) {
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset, errOffset := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if errLimit != nil || errOffset != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Given non-integer value"})
		return
	}

	rows, err := Db.Query(`SELECT id, name, description, mass_grams FROM item LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	defer rows.Close()
	var items []Item = make([]Item, 0)

	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description,
			&item.MassGrams); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
			return
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func getItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("item_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Given non-integer value"})
		return
	}

	var item Item
	row := Db.QueryRow(`SELECT id, name, description, mass_grams FROM item WHERE id=$1`, id)

	errScan := row.Scan(&item.ID, &item.Name, &item.Description, &item.MassGrams)

	switch errScan {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"detail": "Item not found"})
		return
	case nil:
		c.JSON(http.StatusOK, item)
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"detail": errScan.Error()})
		return
	}

}

func postItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	insertSql := `INSERT INTO item(name, description, mass_grams) VALUES ($1, $2, $3) RETURNING id`

	row := Db.QueryRow(insertSql, newItem.Name, newItem.Description, newItem.MassGrams)

	errScan := row.Scan(&newItem.ID)

	switch errScan {
	case sql.ErrNoRows:
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Unable to create item."})
		return
	case nil:
		c.JSON(http.StatusCreated, newItem)
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"detail": errScan.Error()})
		return
	}
}

func connectDb() {
	connStr := "host=localhost port=5432 user=benchapi password=benchapi dbname=items sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	Db = db

	if err := Db.Ping(); err != nil {
		panic(err)
	}
}
