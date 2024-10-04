package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MenuItem struct {
	name     string
	price    float64
	quantity int
}

type Order struct {
	items []MenuItem
	mu    sync.Mutex
}

type MenuOperations interface {
	addItem(item MenuItem)
	displayOrder()
}

func (o *Order) addItem(item MenuItem) {
	o.mu.Lock()
	defer o.mu.Unlock()

	for i, existingItem := range o.items {
		if existingItem.name == item.name {
			o.items[i].quantity += item.quantity
			return
		}
	}
	o.items = append(o.items, item)
}

func (o *Order) displayOrder() {
	o.mu.Lock()
	defer o.mu.Unlock()
	fmt.Println("\n--- Pesanan Anda ---")
	totalPrice := 0.0
	for _, item := range o.items {
		subtotal := float64(item.quantity) * item.price
		fmt.Printf("Menu: %s, Jumlah: %d, Harga Satuan: %.2f, Subtotal: %.2f\n", item.name, item.quantity, item.price, subtotal)
		totalPrice += subtotal
	}
	fmt.Printf("Total Harga: %.2f\n", totalPrice)
}

var menu = []MenuItem{
	{"nasi goreng", 20000, 0},
	{"mie goreng", 15000, 0},
	{"ayam bakar", 25000, 0},
	{"teh manis", 5000, 0},
	{"jus jeruk", 10000, 0},
}

func main() {
	var wg sync.WaitGroup
	order := &Order{}
	reader := bufio.NewReader(os.Stdin)

	printWelcomeMessage()

	for {
		menuName := getMenuInput(reader)
		if strings.ToLower(menuName) == "selesai" {
			break
		}

		validationChannel := make(chan bool)
		wg.Add(1)
		go func(menuName string) {
			defer wg.Done()
			valid := isValidMenu(menuName)
			validationChannel <- valid
		}(menuName)

		isValid := <-validationChannel
		if !isValid {
			fmt.Println("Menu tidak ditemukan, silakan masukkan menu yang tersedia.")
			continue
		}

		quantity := getQuantityInput(reader)
		if quantity == -1 {
			fmt.Println("Input jumlah tidak valid, silakan masukkan angka yang benar.")
			continue
		}

		orderItem := createOrderItem(menuName, quantity)
		wg.Add(1)
		go func(item MenuItem) {
			defer wg.Done()
			order.addItem(item)
		}(orderItem)

		// Menghitung total pesanan menggunakan goroutine lain
		sumChannel := make(chan float64)
		wg.Add(1)
		go func(order *Order) {
			defer wg.Done()
			sumChannel <- calculateOrderTotal(order)
		}(order)

		totalPrice := <-sumChannel
		fmt.Printf("Total sementara: %.2f\n", totalPrice)
	}

	wg.Wait()

	defer fmt.Println("Program selesai")
	order.displayOrder()

	displayEncodedOrderInfo(order)

	processOrdersWithTimeout(order, &wg)
	wg.Wait()
}

func printWelcomeMessage() {
	fmt.Println("Selamat Datang di Restoran Sederhana!")
	fmt.Println("Menu yang Tersedia:")
	for _, item := range menu {
		fmt.Printf("- %s: %.2f\n", strings.Title(item.name), item.price)
	}
}

func getMenuInput(reader *bufio.Reader) string {
	fmt.Print("\nMasukkan nama menu (ketik 'selesai' untuk menyelesaikan pesanan): ")
	menuName, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(menuName))
}

func isValidMenu(menuName string) bool {
	menuName = strings.ToLower(menuName)
	for _, item := range menu {
		if menuName == item.name {
			return true
		}
	}
	return false
}

func getQuantityInput(reader *bufio.Reader) int {
	fmt.Print("Masukkan jumlah pesanan: ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = strings.TrimSpace(quantityStr)

	matched, _ := regexp.MatchString(`^\d+$`, quantityStr)
	if !matched {
		return -1
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return -1
	}
	return quantity
}

func createOrderItem(menuName string, quantity int) MenuItem {
	menuName = strings.ToLower(menuName)
	for _, item := range menu {
		if item.name == menuName {
			return MenuItem{name: menuName, price: item.price, quantity: quantity}
		}
	}
	return MenuItem{}
}

func calculateOrderTotal(order *Order) float64 {
	order.mu.Lock()
	defer order.mu.Unlock()
	totalPrice := 0.0
	for _, item := range order.items {
		totalPrice += float64(item.quantity) * item.price
	}
	return totalPrice
}

func displayEncodedOrderInfo(order *Order) {
	orderInfo := "Detail Pesanan: "
	for _, item := range order.items {
		orderInfo += fmt.Sprintf("%s: %d, ", item.name, item.quantity)
	}
	encodedInfo := base64.StdEncoding.EncodeToString([]byte(orderInfo))
	fmt.Printf("\nInformasi pesanan terenkripsi (base64): %s\n", encodedInfo)
}

func processOrdersWithTimeout(order *Order, wg *sync.WaitGroup) {
	statusChannel := make(chan string, len(order.items))
	for _, item := range order.items {
		wg.Add(1)
		go func(item MenuItem) {
			defer wg.Done()
			select {
			case <-time.After(3 * time.Second):
				statusChannel <- fmt.Sprintf("Pesanan %s gagal diproses karena timeout.", item.name)
			default:
				time.Sleep(2 * time.Second)
				statusChannel <- fmt.Sprintf("Pesanan %s telah selesai diproses.", item.name)
			}
		}(item)
	}

	go func() {
		wg.Wait()
		close(statusChannel)
	}()

	for status := range statusChannel {
		fmt.Println(status)
	}
}
