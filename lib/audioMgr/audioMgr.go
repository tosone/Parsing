package audioMgr

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	uuid "github.com/satori/go.uuid"
	"smartconn.cc/liugen/audio"
)

func init() {
	audio.Startup()
}

// Play 播放一个音频
func Play(path string, startAt int64) {

	fmt.Println("Playing a audio: " + path)
	channel, err := audio.PlayBGM(path, startAt)
	if err != nil {
		log.Println(err)
	}
	<-channel
}

// Resume 恢复播放音频
func Resume() {
	audio.ResumeBGM()
}

// Pause 停止音频
func Pause() {
	fmt.Println("Pausing a audio.")
	audio.PauseBGM()
}

// Stop 停止音频
func Stop() {
	fmt.Println("Stopping a audio.")
	audio.StopBGM()
}

// Break 播放一个音频
func Break(path string) {
	fmt.Println("Breaking a audio: " + path)
	channel, err := audio.PlaySE(path)
	if err != nil {
		log.Println(err)
	}
	<-channel
}

// IsBGMPlaying 是否在播放
func IsBGMPlaying() bool {
	return audio.IsBGMPlaying()
}

// IsBGMPaused 是否停止播放
func IsBGMPaused() bool {
	return audio.IsBGMPaused()
}

// RecordToPath 输入存储路径，返回录音文件名
func RecordToPath(Path string) (fileName string, err error) {
	fileName = uuid.NewV4().String() + ".pcm"
	filePath := path.Join(Path, fileName)

	fp, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	recorder, err := audio.Record()
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(fp, recorder); err != nil {
		fmt.Print("this is err: ", err)
		return "", err
	}

	return fileName, nil
}

// RecordToText 输入最大时长，输出录音识别文字
func RecordToText(max int) (text string, err error) {
	// fmt.Println("start record")
	// recorder, err := audio.Record()
	// fmt.Println("after start record")
	// if err != nil {
	// 	return "", err
	// }

	// text, _, err = STXDetect(recorder, max)
	// if err != nil {
	// 	log.Println(err)
	// 	return "", err
	// }
	// audio.StopRecord()
	// fmt.Println("record over")
	// return text, nil

	time.Sleep(time.Second)
	res := []string{
		"",
		"不好",
	}
	rand.Seed(time.Now().UnixNano())
	return res[rand.Intn(len(res))], nil
}

// StopRecord 停止录音
func StopRecord() {
	audio.StopRecord()
	return
}

// STXDetect 语音识别
func STXDetect(reader io.Reader, max int) (text string, status string, err error) {
	time.Sleep(time.Second)
	res := []string{
		"",
		"不好",
	}
	rand.Seed(time.Now().UnixNano())
	return res[rand.Intn(len(res))], "voice", nil
}
