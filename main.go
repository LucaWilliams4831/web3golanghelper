package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
	

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	pancakeFactory "github.com/LucaWilliams4831/web3golanghelper/contracts/IPancakeFactory"
	pancakePair "github.com/LucaWilliams4831/web3golanghelper/contracts/IPancakePair"
	"github.com/LucaWilliams4831/web3golanghelper/web3helper"

	"database/sql"
    "encoding/json"
    "net/http"
    _ "github.com/lib/pq"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

type Reserve struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}
type Bot struct {
    ID         int64     `json:"id"`
    PrivKey    string    `json:"privkey"`
    Title      string    `json:"title"`
    NetworkID  string    `json:"network_id"`
    Volume     float64   `json:"volume" db:"volumn"`
    Fees       float64  `json:"fees"`
    TotalTx    int64       `json:"total_tx" db:"total_tx"`
    CurTx      int64       `json:"cur_tx" db:"cur_tx"`
    TimeMin    int64     `json:"time_min"`
    TimeMax    int64     `json:"time_max"`
    CreateTime string `json:"create_time"`
    DelFlag    bool      `json:"del_flag"`
    Status     bool      `json:"status"`
}
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "bot_data"
)


func main() {
    // buy()
    db, err := createDBConnection()
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    createTable(db)

    router := mux.NewRouter()

    // Define your routes
    router.HandleFunc("/", handleRequest).Methods("GET")
    router.HandleFunc("/", handleRequest).Methods("POST")
    router.HandleFunc("/{id}", handleRequest).Methods("PUT")
    router.HandleFunc("/{id}", handleRequest).Methods("DELETE")

    // Enable CORS using the handlers package
    headers := handlers.AllowedHeaders([]string{"Content-Type"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"})

    // http.HandleFunc("/", handleRequest)
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))

	// read .env variables
	
}
func buy() {
	RPC_URL, WS_URL, WETH_ADDRESS, FACTORY_ADDRESS, TOKEN_ADDRESS, PK, BUY_AMOUNT, ROUTER_ADDRESS, GAS_MULTIPLIER := readEnvVariables()

	web3GolangHelper := initWeb3(RPC_URL, WS_URL)
	fromAddress := GeneratePublicAddressFromPrivateKey(PK)

	// convert buy amount to float

	// infinite loop
	for {
		// get pair address
		lpPairAddress := getPair(web3GolangHelper, WETH_ADDRESS, FACTORY_ADDRESS, TOKEN_ADDRESS)
		fmt.Println("LP Pair Address: " + lpPairAddress)

		if lpPairAddress != "0x0000000000000000000000000000000000000000" {
			reserves := getReserves(web3GolangHelper, lpPairAddress)

			fmt.Println("Reserve0: " + reserves.Reserve0.String())
			fmt.Println("Reserve1: " + reserves.Reserve1.String())

			// check if reserves is greater than 0
			if reserves.Reserve0.Cmp(big.NewInt(0)) > 0 && reserves.Reserve1.Cmp(big.NewInt(0)) > 0 {
				buyAmount, err := strconv.ParseFloat(BUY_AMOUNT, 32)
				if err != nil {
					fmt.Println(err)
				}
                fmt.Println(buyAmount)
				fmt.Println(web3GolangHelper.GetEthBalance(fromAddress))
				// web3GolangHelper.Buy(ROUTER_ADDRESS, WETH_ADDRESS, PK, fromAddress, TOKEN_ADDRESS, buyAmount, GAS_MULTIPLIER)
				// time.Sleep(10 * time.Millisecond)
				//  web3GolangHelper.Sell(ROUTER_ADDRESS, WETH_ADDRESS, PK, fromAddress, TOKEN_ADDRESS, buyAmount, GAS_MULTIPLIER)
				 web3GolangHelper.Sell(ROUTER_ADDRESS, WETH_ADDRESS, PK, fromAddress, TOKEN_ADDRESS, 10000000000, GAS_MULTIPLIER)
				os.Exit(0)
			}
		}
		// sleep 1 second
		time.Sleep(1 * time.Millisecond)
	}
}

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// function for read .env variables
func readEnvVariables() (string, string, string, string, string, string, string, string, string) {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	RPC_URL := os.Getenv("RPC_URL")
	WS_URL := os.Getenv("WS_URL")
	WETH_ADDRESS := os.Getenv("WETH_ADDRESS")
	FACTORY_ADDRESS := os.Getenv("FACTORY_ADDRESS")
	TOKEN_ADDRESS := os.Getenv("TOKEN_ADDRESS")
	PK := os.Getenv("PK")
	BUY_AMOUNT := os.Getenv("BUY_AMOUNT")
	ROUTER_ADDRESS := os.Getenv("ROUTER_ADDRESS")
	GAS_MULTIPLIER := os.Getenv("GAS_MULTIPLIER")

	return RPC_URL, WS_URL, WETH_ADDRESS, FACTORY_ADDRESS, TOKEN_ADDRESS, PK, BUY_AMOUNT, ROUTER_ADDRESS, GAS_MULTIPLIER
}

func initWeb3(rpcUrl, wsUrl string) *web3helper.Web3GolangHelper {
	web3GolangHelper := web3helper.NewWeb3GolangHelper(rpcUrl, wsUrl)

	chainID, err := web3GolangHelper.HttpClient().NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Chain Id: " + chainID.String())
	return web3GolangHelper
}

func getReserves(web3GolangHelper *web3helper.Web3GolangHelper, pairAddress string) Reserve {

	pairInstance, instanceErr := pancakePair.NewPancake(common.HexToAddress(pairAddress), web3GolangHelper.HttpClient())
	if instanceErr != nil {
		fmt.Println(instanceErr)
	}

	reserves, getReservesErr := pairInstance.GetReserves(nil)
	if getReservesErr != nil {
		fmt.Println(getReservesErr)
	}

	return reserves
}

func getPair(web3GolangHelper *web3helper.Web3GolangHelper, wethAddress, factoryAddress, tokenAddress string) string {

	factoryInstance, instanceErr := pancakeFactory.NewPancake(common.HexToAddress(factoryAddress), web3GolangHelper.HttpClient())
	if instanceErr != nil {
		fmt.Println(instanceErr)
	}

	lpPairAddress, getPairErr := factoryInstance.GetPair(nil, common.HexToAddress(wethAddress), common.HexToAddress(tokenAddress))
	if getPairErr != nil {
		fmt.Println(getPairErr)
	}

	return lpPairAddress.Hex()

}

func GeneratePublicAddressFromPrivateKey(plainPrivateKey string) common.Address {
	privateKey, err := crypto.HexToECDSA(plainPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
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
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    return sql.Open("postgres", connStr)
}

func createTable(db *sql.DB) {
	
    createTableQuery := `CREATE TABLE IF NOT EXISTS bot_tb (
        id SERIAL PRIMARY KEY,
        privkey TEXT NOT NULL,
        title TEXT NOT NULL,
        network_id TEXT NOT NULL,
        volumn DOUBLE PRECISION DEFAULT 0,
        fees REAL DEFAULT 0,
        total_tx INTEGER NOT NULL,
        cur_tx INTEGER DEFAULT 0,
        time_min INTEGER NOT NULL,
        time_max INTEGER NOT NULL,
        create_time TIMESTAMP NOT NULL,
        del_flag BOOLEAN DEFAULT false,
        status BOOLEAN DEFAULT true
    );`
    
    _, err := db.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }
	print("table created successfully\n")
}
func insertData(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        Title       string `json:"title"`
        TotalTx     string `json:"totalTx"`
        Network     string `json:"network"`
        Fees        string `json:"fees"`
        TimeMin     string `json:"timeMin"`
        TimeMax     string `json:"timeMax"`
        PrivKey     string `json:"privKey"`
    }

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        return
    }

    // Process the payload as needed
    fmt.Printf("Received payload: %+v\n", payload)

    db, err := createDBConnection()
    if err != nil {
        http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
        log.Fatal(err)
        return
    }
    defer db.Close()

    // Convert string values to their respective data types
    txNum, err := strconv.ParseInt(payload.TotalTx, 10, 64)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }
    fees, err := strconv.ParseFloat(payload.Fees, 64)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }
    timeMin, err := strconv.ParseInt(payload.TimeMin, 10, 64)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }
    timeMax, err := strconv.ParseInt(payload.TimeMax, 10, 64)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }

    bot := Bot{
        PrivKey:      payload.PrivKey,
        Title:        payload.Title,
        NetworkID:    payload.Network,
        Volume:       0.0,
        TotalTx:      txNum,
        Fees:         fees,
        TimeMin:      timeMin,
        TimeMax:      timeMax,
        CreateTime:   getCurrentTime(),
    }
    

    success := insertBot(db, bot)
    if !success {
        http.Error(w, "Failed to insert to database", http.StatusInternalServerError)
        return
    }


    responsePayload := map[string]string{
        "message": "Success",
    }
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
        http.Error(w, "Failed to send response", http.StatusInternalServerError)
        return
    }

    // w.WriteHeader(http.StatusCreated)
    // fmt.Fprint(w, "201 Created")
}


func insertBot(db *sql.DB, bot Bot) bool {
    query := `INSERT INTO bot_tb (privkey, title, network_id, volumn, fees, total_tx, cur_tx, time_min, time_max, create_time, del_flag, status)
              VALUES ($1, $2, $3, 0.0, $4, $5, 0, $6, $7, $8, false, true)`
        
    _, err := db.Exec(query, bot.PrivKey, bot.Title, bot.NetworkID, bot.Fees, bot.TotalTx, bot.TimeMin, bot.TimeMax, bot.CreateTime)
    if err != nil {
        fmt.Println("insert error: ", err)
        return false
    }
    print("inserted successfully\n")

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

    // success := updateBot(db, bot.ID, bot.Title)
    // if !success {
    //     w.WriteHeader(http.StatusInternalServerError)
    //     fmt.Fprint(w, "500 Internal Server Error")
    //     return
    // }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "200 OK")
}

func updateBot(db *sql.DB, id int, newTitle string) bool {
    updateDataQuery := `UPDATE bot_tb SET title = $1 WHERE id = $2;`

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
    // w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func queryBots(db *sql.DB) []Bot {
    queryDataQuery := `SELECT id, privkey, title, network_id, volumn, fees, total_tx, cur_tx, time_min, time_max, create_time, del_flag, status FROM bot_tb ORDER BY create_time DESC;`
    print(queryDataQuery)
    rows, err := db.Query(queryDataQuery)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    bots := []Bot{}
    for rows.Next() {
        var bot Bot
        err := rows.Scan(&bot.ID, &bot.PrivKey, &bot.Title, &bot.NetworkID, &bot.Volume, &bot.Fees, &bot.TotalTx, &bot.CurTx, &bot.TimeMin, &bot.TimeMax, &bot.CreateTime, &bot.DelFlag, &bot.Status)
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
    deleteDataQuery := `DELETE FROM bot_tb WHERE id = $1;`

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
