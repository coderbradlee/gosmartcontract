package highlimit

import (
	"math/big"
	"regexp"
	"strconv"
	"time"
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

var Ether = math.BigPow(10, 18)
var Shannon = math.BigPow(10, 9)

var pow256 = math.BigPow(2, 256)
var addressPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
var zeroHash = regexp.MustCompile("^0?x?0+$")

func IsValidHexAddress(s string) bool {
	if IsZeroHash(s) || !addressPattern.MatchString(s) {
		return false
	}
	return true
}

func IsZeroHash(s string) bool {
	return zeroHash.MatchString(s)
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	diff1 := new(big.Int).Div(pow256, difficulty)
	return string(common.ToHex(diff1.Bytes()))
}

func TargetHexToDiff(targetHex string) *big.Int {
	targetBytes := common.FromHex(targetHex)
	return new(big.Int).Div(pow256, new(big.Int).SetBytes(targetBytes))
}

func ToHex(n int64) string {
	return "0x0" + strconv.FormatInt(n, 16)
}

func FormatReward(reward *big.Int) string {
	return reward.String()
}

func FormatRatReward(reward *big.Rat) string {
	wei := new(big.Rat).SetInt(Ether)
	reward = reward.Quo(reward, wei)
	return reward.FloatString(8)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

func String2Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)
	return n
}
func SmsSend(mobile []string, content string) {
	var smsSign = "【91Pool】"
	for _, v := range mobile {
	str := "username=xujing&password=xjMall!@#136xj&mobile=" + v + "&content=" + content + smsSign + "&xh="
	var result string
	var err error
	for i := 0; i < 10; i++ {
	result, err = sendMessage("http://114.215.196.145/sendSmsApi", str)
	if err != nil {
		log.Printf("次发送短信失败，五秒钟后重新发送", i+1, err, result)
	time.Sleep(time.Second * 5)
	continue
	}
	break
	}
	log.Printf("发送短信:" + v + "[" + content + "],返回结果:" + result)
	}
}
func Sendmessage(content string){
	//	
	mobile:=[]string{"18930126300","18221300379"}
    SmsSend(mobile,content)	
}
func sendMessage(uri, reqBody string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
	return "Create request error", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cache-control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
	return "Send message error", err
	}
	defer resp.Body.Close()
	respStr, _ := ioutil.ReadAll(resp.Body)
	return string(respStr), nil
}