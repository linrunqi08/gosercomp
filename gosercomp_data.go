package main

import (
	"time"

	flatbuffers "github.com/google/flatbuffers/go"

	fb "github.com/smallnest/gosercomp/fbs/LogGroup"
	"github.com/smallnest/gosercomp/messagepack"
	massagepack "github.com/smallnest/gosercomp/messagepack"
	protocol "github.com/smallnest/gosercomp/protocol"
)

func CreateLogs(kvs ...string) *protocol.Log {
	var slsLog protocol.Log
	for i := 0; i < len(kvs)-1; i += 2 {
		cont := &protocol.Log_Content{Key: kvs[i], Value: kvs[i+1]}
		slsLog.Contents = append(slsLog.Contents, cont)
	}
	slsLog.Time = uint32(time.Now().Unix())
	return &slsLog
}

func CreateLogGroup() *protocol.LogGroup {
	logGroup := &protocol.LogGroup{}
	for i := 0; i < 10; i++ {
		logGroup.Logs = append(logGroup.Logs, CreateLogs("kkkkkkkkeeeeyyyyyyy0", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy1", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy2", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy3", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy4", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy5", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"))
	}
	logGroup.Category = "testLogstore"
	logGroup.Source = "testSource"
	logGroup.Topic = "testTopic"
	return logGroup
}

func CreateMPLogs(kvs ...string) *massagepack.Log {
	var slsLog massagepack.Log
	for i := 0; i < len(kvs)-1; i += 2 {
		cont := &massagepack.LogContent{Key: kvs[i], Value: kvs[i+1]}
		slsLog.Contents = append(slsLog.Contents, cont)
	}
	slsLog.Time = uint32(time.Now().Unix())
	return &slsLog
}

func CreateMpLogGroup() *massagepack.LogGroup {
	logGroup := &massagepack.LogGroup{}
	for i := 0; i < 10; i++ {
		logGroup.Logs = append(logGroup.Logs, CreateMPLogs("kkkkkkkkeeeeyyyyyyy0", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy1", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy2", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy3", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy4", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"kkkkkkkkeeeeyyyyyyy5", "value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"))
	}
	logGroup.Category = "testLogstore"
	logGroup.Source = "testSource"
	logGroup.Topic = "testTopic"
	return logGroup
}

func CreateFpLogGroup() []byte {
	builder := flatbuffers.NewBuilder(0)

	logs := make([]flatbuffers.UOffsetT, 0)
	for j := 0; j < 10; j++ {
		logContents := make([]flatbuffers.UOffsetT, 0)
		for i := 0; i < 6; i++ {
			keyh := builder.CreateString("kkkkkkkkeeeeyyyyyyy0")
			valueh := builder.CreateString("value1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
			fb.LogContentStart(builder)
			fb.LogContentAddKey(builder, keyh)
			fb.LogContentAddValue(builder, valueh)
			content := fb.LogContentEnd(builder)
			logContents = append(logContents, content)
		}
		fb.LogStartContentsVector(builder, 6)
		for _, content := range logContents {
			builder.PrependUOffsetT(content)
		}
		rLogContents := builder.EndVector(6)
		fb.LogStart(builder)
		fb.LogAddTime(builder, time.Now().Unix())
		fb.LogAddContents(builder, rLogContents)
		log := fb.LogEnd(builder)
		logs = append(logs, log)
	}
	fb.LogGroupStartLogsVector(builder, 10)
	for _, log := range logs {
		builder.PrependUOffsetT(log)
	}
	rlogs := builder.EndVector(10)

	categoryh := builder.CreateString("testLogstore")
	sourceh := builder.CreateString("testSource")
	topich := builder.CreateString("testTopic")

	fb.LogGroupStart(builder)
	fb.LogGroupAddLogs(builder, rlogs)

	fb.LogGroupAddCategory(builder, categoryh)
	fb.LogGroupAddSource(builder, sourceh)
	fb.LogGroupAddTopic(builder, topich)
	logGroup := fb.LogGroupEnd(builder)
	builder.Finish(logGroup)
	buf := builder.FinishedBytes()
	return buf
}

func DecodeToFpLogGroup(buf []byte) *massagepack.LogGroup {

	// buf, err := ioutil.ReadFile(filename)
	// if err != nil {
	//     panic(err)
	// }
	//传入二进制数据
	b := fb.GetRootAsLogGroup(buf, 0)
	len := b.LogsLength()
	logGroup := &messagepack.LogGroup{}

	//count := 0
	for i := 0; i < len; i++ {
		tx_f := new(fb.Log)
		tx := new(massagepack.Log)
		if b.Logs(tx_f, i) {

			tx.Time = uint32(tx_f.Time())
			contentLen := tx_f.ContentsLength()

			cont_f := new(fb.LogContent)
			//cont := new(massagepack.LogContent)
			for j := 0; j < contentLen; j++ {
				if tx_f.Contents(cont_f, j) {
					_ = string(cont_f.Key())
					_ = string(cont_f.Value())

					//key := string(cont_f.Key())
					//value := string(cont_f.Value())
					//fmt.Println("Key:", key)
					//fmt.Println("Value:", value)
					//fmt.Println(count)
					//count++

				}
			}
			//tx.Contents = append(tx.Contents, cont)
		}
		//logGroup.Logs = append(logGroup.Logs, tx)
		//append(block.Children, *tx)
		//fmt.Println("child:", tx)
	}
	return logGroup
}

/*
var group = model.ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var avroGroup = model.AvroColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var zgroup = model.ZColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var egroup = model.EColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var fgroup = model.FColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var thriftColorGroup = model.ThriftColorGroup{
	ID:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var protobufGroup = model.ProtoColorGroup{
	Id:     proto.Int32(17),
	Name:   proto.String("Reds"),
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var gogoProtobufGroup = model.GogoProtoColorGroup{
	Id:     proto.Int32(1),
	Name:   proto.String("Reds"),
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var colferGroup = model.ColferColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var avroSchema = `{"namespace": "gosercomp",
"type": "record",
"name": "ColorGroup",
"fields": [
	 {"name": "id", "type": "int"},
	 {"name": "name",  "type": "string"},
	 {"name": "colors", "type": {"type": "array", "items": "string"}}
]
}`

var rlpgroup = model.RlpColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var gojayGroup = &model.GojayColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var thrfitIterGroup = model.ThriftIterColorGroup{
	ID:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var thrfitIterGroupDynamic = general.Struct{
	1: int32(1),
	2: string("Reds"),
	// 3: general.List{string("Crimson"), string("Red"), string("Ruby"), string("Maroon")},
	3: general.List{"Crimson", "Red", "Ruby", "Maroon"},
}
*/
