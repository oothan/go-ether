package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	keys := "2e36f270c48945bfade2661724b52075|bed4024203054ca3a944edfe167ae507|b2edac3c231f4ef08c0d51b5c02e4cfc|1f1990dd4cc7431f9da872a1e5732093|aff096ef309945c2841a5bd4c555a3c6|92515e388d334031baa5f3d69bbe4e52|8d86a5f4167347008161a347986354a4|22b746b8a7424331bedff448db7b804a|d23e3f58177a49dd87f4381cc75ab166|7ce613fd939d4089a79ae92b0f82364c|74fe3ad0a7b04dfcaa8e371e18351e07|60bdff8105964242ac4adac351482fec|698a83485dd24309b21e049a90632859|4b126811b51045078387e5099d5a0101|8eb9cd7e28d64150bc30ed6af73edaea|518f6c423c154dcbbbec051091570503|3db5794896ad4a55adab2678b7f4dec4|505959cb48b9462aa931da4c2afd8385|b92eb28803a342519d37c7536e83b2f4|bac1c11252bd4d3aa1aea44c8f99a4f2"
	keySplit := strings.Split(keys, "|")
	for i, k := range keySplit {
		fmt.Println(i, " ", k)
	}
	fmt.Println("Length : ", len(keySplit))



	ticker := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-ticker.C:
			min := 0
			max := len(keySplit) - 1

			rand.Seed(time.Now().UnixNano())
			tick := rand.Intn(max - min + 1) + min
			fmt.Println(tick)
			fmt.Println(keySplit[tick])
		}
	}
}
