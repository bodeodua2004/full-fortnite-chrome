package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.2" )

func qtebctVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMAMYrBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojR6TWF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWD4QzNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rFAhGkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydnoQMK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAHYajJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIzZFT24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xoM29ldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5jWYzAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBnGgu1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daLOM5KMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lb9FpR18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaRlnv63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGr129p1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1GpgQPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jz5VPE5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gwkp516zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KcOGrzK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1eYr8xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1A3vlQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjHkspdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxzZK6UWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTevUEyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jk3vXvSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5FgkcXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPnSV5n3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kniHWJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TndhmvvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDXZNWBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMVanrzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbU6ofU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HP8AEd37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0B2SCvltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KqzSMzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2bGbld6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sVa5dR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBzzFFBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7HRsuVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVTtO0sVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6H89A1vPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1V3JMQXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XB5994qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SO5DQQQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvNmmGG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhgv2ksOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLh3KuK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Er5PB6ArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmvFGAgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAjOJ4UdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwbTCd47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90z3CXYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kMlfzlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8QFYc7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kvI8W1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnorA6sxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PFGabaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvqCBkghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5cN4y9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRAxlrHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRxKFmMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLKIAbbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K57CnREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWnWc1jDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pp47sNoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQV1KuAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVvijv9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrzkI20NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1MK1WSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udBDWt1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNFjGjNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgDtSQltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSgKxOM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKghAb4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeyBVxRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3AtOA2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTWcPHrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uB04AxfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1OjIM01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFGq8O4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func np75yjllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfn069qJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ohmzz6DLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyiQaRPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYmaiajHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwwP0jZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsi2Ldd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDd8eEgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IU3Vd1QWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6OcrJhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njYVSiOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsLUo31PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJCKzl5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AodLkVNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTN7JvQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTaPveOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFuYRYw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJ5eB44SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zr3tnklAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObK9GSSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UkiW63EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbaIe76sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmgSMK50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUfTmNvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uoKYmIsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qf3cFfF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOkw1YmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUF9TDXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxscfqJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tH0tBTjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhqMROlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NbMAbDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLLRJp38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gItlBWMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qZFCzouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuMfRBYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40qin0i7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvZewuOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVhXnLJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKoP9HDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkQQMb7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QgSt0vqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLUjjjuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEdouHEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VIrkr7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOEZ0xoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2H1wCefOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfS2AMV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grvuUV1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCiPkmLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Pb2I96iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfNEN81YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 228J34RnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjmyVuyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOI04vaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnIttLAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebpB9NDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcKNB0MwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7qjic9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQcJJrjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfk6JuGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FscWs1EIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBPte9xLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqEzeQ5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Av2gxqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JfNHCx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csKCNpjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmu1Hd6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LbhouuhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mb2jXkz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQdpkDVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2skruETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ClTKo3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQWWHOGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDccT9twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EYhTSrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixAnbzOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmyjITO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLVaKiGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyW67R4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0brwpjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Smf9BxvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2WXXE5QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6hBu71zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJIabwOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGGbnhMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5Vw6EGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzupLeCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q38GteiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsZNurrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cy1IKr1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhNmQkTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Q5naZyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wn2tURERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASE35Eg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zlvNJQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xs7cahVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83SV5vSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 233X5rn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvAlx90zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJUelF5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BNJb2QHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJljL7iUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfAoXbriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eD96l3N9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvexFVSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJxd941UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmoI1ojpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJZNg3DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZveSjS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHsenFYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVJAvO8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWGEVBnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OuUlwpK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgzZC0U6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcPSG1FSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t79Xk0daWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CU86Aj0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwJCx2GEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KYjViA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBbQKrteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoUgxeMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrq8MKkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjOmsIRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVJ3U2vzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZB3WHNHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDoARRQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKxwnI38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qWJHmXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41kPcPUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSTABF8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GusYaVWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9frOERLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9DENmFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwcZGP5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KE3Yt4NVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jC7IaLUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qH4kiKXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNrHnZFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avuACRE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MnEdKaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXRgUOFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0Uw7HPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7FinFcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvABFmc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3q0Y7dWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dgkAWZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4DcrJFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func coLaMN3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItnEmTM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVClrMLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxxYv36qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iV8izJ89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYQsyOFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c852hZxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twDy79aRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sO8yG91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ro4nnIIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iTsHqKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiUStzhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EF9r0MRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aF8oQqmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkJS8Lj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umtgH8xsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOwnVZQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4dnqRioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh5ZFhPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzR8ZRmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4bYGZwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwEPowq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuL8jcGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M68uZVkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vD7pUFpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umSVsiCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doy2LRgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lX5IKjkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXkpTtLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9mGsKHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGA3T5VsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61x24kXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRZIdKkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpZfnnhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59A2dQFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX4tktAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHzbMhhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFUheRlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ykXUo8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSmCdc0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbRorPGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQRhvNfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c10XYl0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCt3PAiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmeic1t6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQN369rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQHnJ5BXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWwsCxcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4WpegytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7TaCKUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7enjI8xkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8UEDLTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzusDpq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lR9x64CmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBUq3C2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8n40XreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y78qjSLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RX5KHrQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTg6lr88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksrB2UQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcj13NEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F38pWR3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYOgt8pwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zl85b8wJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZy9ugHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PptYFqPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89I3dMflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERQwWXltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlirTFTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMn0bnHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xA39WtNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OrY1g6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUYJc3s8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzgn0j0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XVuPGr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DJTqGm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8QXIigsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qj4pTXbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi2KKNrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FR9vT0ZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouopuEFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73JbEBPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoD51MPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIWCpkK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DFDjYEy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFtlt2GAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLcpCVJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0dxvgwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vYrH3ycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTeOkUiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwd46HVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psNRcwTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBKRnowtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsgW06qHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ykw1fi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgmNZDVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Kmj9GQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4hEREwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8x38oHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gu3K1rOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLy963HAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eRDfS8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwoMgqApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etZn69wbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8p97RIROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PKzLeo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtOR67poWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMhiBrJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9XEPVPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3U39HblxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBgvTzdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZ8clbgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9UwigXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUtELg06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHJqwXBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwV1igugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7u5izioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxQV2O7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YB5zxJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1KYBotQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A712c0gaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUjrmaj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEDYLsNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3CO1DFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpqqMzKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5iqsBUkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JrVqAEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilyZuHA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYmyWqTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1o8eEgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vBlhnDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Elm666A4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfylV4IaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvkeiGaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3vVK4toWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Xn1H34XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saDfYg7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2wgh2KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cuae59FbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R11hU3w5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNdaWMEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67C5OewtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOzPfcgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHE6mRTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lo3mRuPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EC0dZbJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qNnH2FZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egC1Jy8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5YyFRtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsXDYBzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2ZQH6T0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOROo8QPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tvk08MSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVcmBjG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIqMJSZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvHEl4OWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mdir47kBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A56kB1liWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7s29TBRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7x0kQVuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E86askHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rc66PiMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jx7EbDOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyfxFdxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBmCaN12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTbHPxjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylFQCgz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZ77fLbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wjsIIHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3Sl4JFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwZDpn0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeMmmKqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faipUBb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFCllxiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UU68TZhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSpXjOtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiwbFRFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDuvfhUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQUbUTSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oRKghtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nt6COhZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IeKl7vcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fF7rmevLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLrH9sTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZwT6Q7z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvPHoVUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4V7pAxEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLdmJ9wKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agDULUETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIUSBa9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suXt9Pi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGn1EIEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LrY1uu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OfaTPXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BPItBWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QH7D6PnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2DJojvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpFZHmzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjuGJ4YUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KfCKPGu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65eTRhzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnMf0UrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bS7CisRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROP6iAhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlsoW430Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqdYUD7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoqCzwpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ByIT2AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hrnq1zMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OL612X2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cR20N4QBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwTFuUQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Z0SJaCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eU6fozmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pb46Wk1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9Jmb2LrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soRPxtBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDdPbTy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqarpVzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H74i3lULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLzL5YzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZmGal1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS3FuO74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFwNA4vjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nP6cTqiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNwZ5gldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfrnft0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exTdDJGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPeOdFZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTjUO3fJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A74dBbFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfRYlcp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUZLnGa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DiUFMSgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpXCdWOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBuIbJIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ksImGqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qG1gURAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2eVlYeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbpcsrAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVOl7VdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9foBDzxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpy0Um57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3B9Y8KSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5ZAupCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9anooDWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLcTCCmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p07ZJQwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbvLcvtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COoVIt5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKR3goR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4sSmXsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKgf0M31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIYDne4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rB4sNfofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQAYnb2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2AwTmK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWvfcSzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTKgxKpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWJxGwxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZBxLxi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kB3DUMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ug0WNxvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9J7mo4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1BEXHswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOcZ8t3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhbqFlgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ViHrvDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGlF440MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2bmpKIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoFV46uUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2WkdAf3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54wFln02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LorVTQ2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1LmbHbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRk7LPkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgawDB97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UD2ohL1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX8fwCwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfI0KCDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJmirszxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rE1UaXcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8q4yIO0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2L1MnXHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ftp5ZrbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9cugMpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zv7SZThgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBOr5PxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8ulKpMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMpNl7oYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tmu8MEwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3C6n3iAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS6qIRYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rowynepbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXRQXlV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUcnc6bbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXue4PptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpiJydDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hd2MSEaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKnTRWZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZgKXeT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDAorGEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFSs6VaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zk4TSaAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LN9Ojw8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kj3uhB3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kATNmvKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLlCrZRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func naLAHCRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFIuT97eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXnAk4z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuIi1fsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hff1u4EfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcQFf3WdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6e5nx4WxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCMS9IhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UlCKuTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGVZnvMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXypUWmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBOgAGuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgEdV3ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70ImoP8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdNz9mq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggGy0GWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbXjEU0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGE73CApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKPJAjakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7efDGVc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADUzkEi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEmWO8o1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqrXqVIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMuO3BvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE4s5HGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBTyyTkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ri70SwRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wNeOHfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8DY5VSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqFEka5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePymN2sWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xi0K4HjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDGc2pNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhcMKcuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEcc7JT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adoAItduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UmtgV4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeRAw2fPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9IcwgtMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lE9TSNdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOuDKTZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qqi2tNHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HkfR6c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dY5IbEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJpqwoZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0Y0VkKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Spc5YS65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxRP3FNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDYrOGZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXFaQBvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ28eLpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MmZtnFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2QYFrhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2N1F8galWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gu54kTaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hGS5CyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4utO4S4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUZUT0RtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1yHp9A8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CglGCS0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znmwnUcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlJ3iWyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJpxfAdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyNUAutQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wepdn9lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tah4KCgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzrFCXzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4xHvZHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfX85KToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMM1boZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaPdy70EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLc80vRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x48JHL6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaw7wrg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1ksgme4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uj6Qov04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKGHWz8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mr1IQcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5pFJmPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVlhIJnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAurkbDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0tswkalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZcMHdYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iI23laK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFP5XesKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func keZcD6f0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElhH7WDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4C42sxBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbXIFJJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MmY2powFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPDAHR21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E726VQRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mH72FzMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HibPh9jmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pbg7y5JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDJlU0iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CU2KZR11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKFDrwj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8i0THGFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6gg3gjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkgFtGDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sy7G9IgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZkWCp1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ke87bTfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qzapocnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dz68b4JNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1jtkL5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gw1gaRinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O16CNi3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6yyXLYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pphCYRaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86ViAlTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJvp1PaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPHpinBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woAihF2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UO3ZPsCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eontZ1wxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RorqkEsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mwu3ykfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KvyYLHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJyc0UaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9k91mNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJlKUrGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9vA5k1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmtK25ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vFTpqQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9sdpko8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HazdxWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izKeRxtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnXy5Tm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dzJ3cXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHl5VY0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCBCXIvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2FvyfWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBY0MCQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7h8wZldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rR2V8MfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5GkI86wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHDBkqo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41DbCk99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aJ5KWX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OMLSsDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NruFkbOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ym1bA8mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NcWYzazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuUpYEpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nA7IrPFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHpiwTQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNk9YljFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZlTvNGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zq7Imi7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ablx84gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsjEye3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Or1cSFcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueRjJiUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUIdEmnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X16h3jxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z61Wi4BKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNk6vemUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOKk4feTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0fWKhaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxnkM8B7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhV8qt0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eX02xHATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4IDiSWzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sa8yfMvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7WOr81kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGFbwl4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Noud7xy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJQ2bPYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YI66XdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KP3yDTh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyAfd2xzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQg3t971Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRmW81WOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8QQkCBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CluwD3mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuFSfciBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBmRJiSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1nIGpVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9VmQj7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Yz8OG6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7aLQTvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHmNfACpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOsyoJnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WacoPJsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JNDeuoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EukSQlZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuzqeYe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyZcZD3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiV5cqSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5RLexVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCE2DJ0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BthNkoRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmixurjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQNftXfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W15NfIgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bu6VHH8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKJt9AwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vT4YekTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQPFkyhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJO70yTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOKvxHXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvTlg4yHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4NqfB3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbfmqgt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ab46jDnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bh6hPzJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VySOimxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3oqnZW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKnJcwADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZ9G0XGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXG4dhiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sU9WsemHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmnMIE7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OOlMbImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhGOELRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSeLqnhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vShmxP5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwZJdDbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9WuBFdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4Blf12KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8GxEU1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wO8TjqBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXETc6VjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwajHXXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DR6rCpItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGkuufnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77vJb3UjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KAxYgBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrC6sHcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS5uONwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DKQjrbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19Pwo5tTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyrE6pbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0wBaQtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB4AK3n0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpLfRE2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WetrCuABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsm4lTYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DdB61vmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gF0wGXzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96EnFRVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpsLe890Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hfL2YhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUhLl6DAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ue7ghAGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rb7SdcdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mp3VR5V2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM3W8MMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObPTkWaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caBZovWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AIBvtKPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFJ2FRVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rroV2xaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Syl2vLVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoMyzsabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWN64AygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93r5IyANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fi0oOhLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXPJuHNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z18zhEMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWFtI4iiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7cvcQatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lm5YHPWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URwbWEAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JAKXHKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWq6DZuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8b7MxelCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7o42VeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4x8Dz6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7FcuC08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NweOw1edWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTm4Ghr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irOVWVdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmsqIKNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2a78AK1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOjHaEnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9SBVLLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rsOKQBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1NWEVFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnoWL2J9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVuut6wtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuVaLhFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQRrYUpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icrjNZooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rvJn9IlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfuFI0KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCgKKiQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kewMxM6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWYSTfWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3ANBYCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZpOxJjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8NwiV6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRYg79ATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZP1GwsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRG318wmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0JnfG78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDP4FyoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6UnSyrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMvcaaWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfpSETZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdXZFUhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6W3B2qTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oboeFBoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYigbfjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VG4zWrBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQyRuV3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GpnKtf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HDQHxKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iHaXiFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MaQB67SVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgkRE5wHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daRMd9GZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgUV0InSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wi3K6163Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8EBFDWT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osarfFI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxmEasJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wiqy7R8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NK2UOzXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rCw53HnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYMzvcX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPotDe6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJUVLFtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOqu3LyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZJCDT2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKe5CGxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XYRKcaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLQAt5jHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqZD30ueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDBjoIBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8meBrHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wkq64PsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

