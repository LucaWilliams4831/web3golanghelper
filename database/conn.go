package database

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    "strconv"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "bot_data"
)

type Bot struct {
    ID            int            `json:"id"`
    PrivKey       string         `json:"privkey"`
    Title         string         `json:"title"`
    NetworkID     string         `json:"network_id"`
    Volume        float64        `json:"volume"`
    Fees          float32        `json:"fees"`
    RandomAction  int            `json:"random_action"`
    TxNum         int            `json:"tx_num"`
    MaxGas        float32        `json:"max_gas"`
    TimeMin       int            `json:"time_min"`
    TimeMax       int            `json:"time_max"`
    CreateTime    string         `json:"create_time"`
    DeleteDateTime sql.NullString `json:"delete_time"`
}

func connect() {
    db, err := createDBConnection()
    if err != nil {
        log.Fatal(err)
    }
    
    createTable(db)

    bot1 := Bot{
        PrivKey:      "mock priv key 1",
        Title:        "bot 1",
        NetworkID:    "network id 1",
        Volume:       0.0,
        Fees:         0.0,
        RandomAction: 0,
        TxNum:        0,
        MaxGas:       0.0,
        TimeMin:      1,
        TimeMax:      10,
        CreateTime:   getCurrentTime(),
    }

    bot2 := Bot{
        PrivKey:      "mock priv key 2",
        Title:        "bot 2",
        NetworkID:    "network id 2",
        Volume:       0.0,
        Fees:         0.0,
        RandomAction: 0,
        TxNum:        0,
        MaxGas:       0.0,
        TimeMin:      1,
        TimeMax:      20,
        CreateTime:   getCurrentTime(),
    }

    insertBot(db, bot1)
    insertBot(db, bot2)
    

    
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        queryData(w)
    case http.MethodPost:
        insertData(w, r)
    case http.MethodPut:
        updateData(w, r)
    case http.MethodDelete:
        deleteData(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprint(w, "405 Method Not Allowed")
    }
}

func createDBConnection() (*sql.DB, error) {
    print("db connected")
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    return sql.Open("postgres", connStr)
}

func createTable(db *sql.DB) {
	
    createTableQuery := `CREATE TABLE IF NOT EXISTS bots_info (
        id SERIAL PRIMARY KEY,
        privkey TEXT NOT NULL,
        title TEXT NOT NULL,
        network_id TEXT NOT NULL,
        volumn DOUBLE PRECISION DEFAULT 0,
        fees REAL DEFAULT 0,
        random_action INTEGER DEFAULT 0,
        tx_num INTEGER DEFAULT 0,
        max_gas REAL DEFAULT 0,
        time_min INTEGER NOT NULL,
        time_max INTEGER NOT NULL,
        create_time TIMESTAMP NOT NULL,
        delete_time TIMESTAMP
    );`
    print(createTableQuery)
    _, err := db.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }
	
}
func insertData(w http.ResponseWriter, r *http.Request) {
    db, err := createDBConnection()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        log.Fatal(err)
        return
    }
    defer db.Close()

    var bot Bot
    err = json.NewDecoder(r.Body).Decode(&bot)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, "400 Bad Request")
        return
    }

    bot.CreateTime = getCurrentTime()

    success := insertBot(db, bot)
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, "201 Created")
}

func insertBot(db *sql.DB, bot Bot) bool {
    insertDataQuery := `INSERT INTO bots_info (privkey, title, network_id, volumn, fees, random_action, tx_num, max_gas, time_min, time_max, create_time) VALUES
        ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
        print(insertDataQuery)
    _, err := db.Exec(insertDataQuery, bot.PrivKey, bot.Title, bot.NetworkID, bot.Volume, bot.Fees, bot.RandomAction, bot.TxNum, bot.MaxGas, bot.TimeMin, bot.TimeMax, bot.CreateTime)
    if err != nil {
        log.Println(err)
        return false
    }

    return true
}

func updateData(w http.ResponseWriter, r *http.Request) {
    db, err := createDBConnection()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        log.Fatal(err)
        return
    }
    defer db.Close()

    var bot Bot
    err = json.NewDecoder(r.Body).Decode(&bot)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, "400 Bad Request")
        return
    }

    success := updateBot(db, bot.ID, bot.Title)
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "200 OK")
}

func updateBot(db *sql.DB, id int, newTitle string) bool {
    updateDataQuery := `UPDATE bots_info SET title = $1 WHERE id = $2;`

    _, err := db.Exec(updateDataQuery, newTitle, id)
    if err != nil {
        log.Println(err)
        return false
    }

    return true
}

func queryData(w http.ResponseWriter) {
    db, err := createDBConnection()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        log.Fatal(err)
        return
    }
    defer db.Close()

    bots := queryBots(db)

    response, err := json.Marshal(bots)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func queryBots(db *sql.DB) []Bot {
    queryDataQuery := `SELECT * FROM bots_info ORDER BY create_time DESC;`
    print(queryDataQuery)
    rows, err := db.Query(queryDataQuery)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    bots := []Bot{}
    for rows.Next() {
        var bot Bot
        err := rows.Scan(&bot.ID, &bot.PrivKey, &bot.Title, &bot.NetworkID, &bot.Volume, &bot.Fees, &bot.RandomAction, &bot.TxNum, &bot.MaxGas, &bot.TimeMin, &bot.TimeMax, &bot.CreateTime, &bot.DeleteDateTime)
        if err != nil {
            log.Fatal(err)
        }
        bots = append(bots, bot)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    return bots
}

func deleteData(w http.ResponseWriter, r *http.Request) {
    db, err := createDBConnection()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        log.Fatal(err)
        return
    }
    defer db.Close()
    print("hello")
    
    if err := r.ParseForm(); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, "400 Bad Request")
    }

    id, err := strconv.Atoi(r.FormValue("ID"))
    if err != nil {
        fmt.Println("Param Error: ", err)
        return
    }
    success := deleteBot(db, id)
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 Internal Server Error")
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "200 OK")
}

func deleteBot(db *sql.DB, id int) bool {
    deleteDataQuery := `DELETE FROM bots_info WHERE id = $1;`

    _, err := db.Exec(deleteDataQuery, id)
    if err != nil {
        log.Println(err)
        return false
    }

    return true
}

func getCurrentTime() string {
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    return currentTime
}
