package hostsnap

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
)
const defaultjson = `{
    "data":
        {
        "mem":
            {
            "meminfo":
                {
                "total":1234567891010
                }
            }
        "disk":
                {
                "usage":[
                     {"total":10000000000}
                     {"total":10000000000}
                     {"total":12345678910}
                     {"total":10000000000}
                     {"total":10000000000}
                    ]
                }
        }
}`
//使用parse函数得到disk
func getDiskFromParse(value *gjson.Result) uint64{
	setter := parseSetter(value,"127.0.0.1","0.0.0.0")
	// 获取bk_disk字段
	cv,ok:= setter["bk_disk"]
	Expect(ok).To(BeTrue())
	// 将结果转化为uint64
	results,ok := cv.(uint64)
	Expect(ok).To(BeTrue())
	return results
}
//自己计算disk
func getDiskFromSelf(value *gjson.Result) uint64{
	var sum uint64
	arr := value.Get("data.disk.usage.#.total").Array()
	for _,v := range arr{
		sum += v.Uint() >> 10 >> 10 >> 10
	}
	return sum
}
//使用parse函数得到mem
func getMemFromParse(value *gjson.Result) uint64{
	setter := parseSetter(value,"127.0.0.1","0.0.0.0")
	// 获取bk_disk字段
	cv,ok:= setter["bk_mem"]
	Expect(ok).To(BeTrue())
	// 将结果转化为uint64
	results,ok := cv.(uint64)
	Expect(ok).To(BeTrue())
	return results
}
//自己计算mem
func getMemFromSelf(value *gjson.Result) uint64{
	return  value.Get("data.mem.meminfo.total").Uint() >> 10 >> 10
}
var _ = Describe("Hostsnap", func() {
	Context("test key-disk", func() {
		It("", func() {
			// 解析为gjson.Result
			gson := gjson.Parse(defaultjson)
			// 求出parse函数得出的结果与理论结果
			shouldout :=getDiskFromSelf(&gson)
			result:= getDiskFromParse(&gson)

			Expect(shouldout).To(Equal(result))

		})
	})
	Context("test key-mem",func(){
		It("", func() {
			// 解析为gjson.Result
			gson := gjson.Parse(defaultjson)
			// 求出parse函数得出的结果与理论结果
			shouldout := getMemFromSelf(&gson)
			result:=getMemFromParse(&gson)

			Expect(shouldout).To(Equal(result))

		})
	})
})
