package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	//"os"
	//"xyz"
	"github.com/davidbarbera/articulate-pocketsphinx-go/xyz"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Ps(jsgf_buffer []byte, audio_buffer []byte, params []string) []Utt {
	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

	xyz.Ps_call(jsgf_buffer, audio_buffer, params, result)

	//Adapting result from coded string to utt struct
	raw := strings.Split(result[0], "**")

	fmt.Printf("%T", raw)
	fields := strings.Split(raw[0], "*")

	fmt.Println(fields)
	hyp := fields[0]
	header := fields[1]

	fmt.Println(hyp)
	fmt.Println(strings.Split(header, ","))
	utts := []Utt{}
	//var utts = make([]Utt, len(fields)-2)

	for i := 0; i < len(fields)-2; i++ {
		parts := strings.Split(fields[2:][i], ",")
		phoneme := parts[0]
		text_start := parts[1]
		text_end := parts[2]
		start, serr := strconv.Atoi(text_start)
		end, eerr := strconv.Atoi(text_end)

		if phoneme != "(NULL)" {
			fmt.Println(phoneme, start, end)
			utts = append(utts, Utt{phoneme, int32(start), int32(end)})

			if serr != nil || eerr != nil {
				fmt.Println(serr, eerr)
			}
		}
	}
	return utts
}

func main() {
	//Data:
	//Filenames
	//filename_jsgf := "./../data/kl_ay_m.jsgf"
	//filename_wav := "./../data/climb1_colin.wav"

	filename_jsgf := "/home/dbarbera/Data/climb/forced_align_5a02788e-3d91-446f-828a-428b7f2d8785_climb1_colin_fixed_trimmed.wav.jsgf"
	filename_wav := "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav"
	//Parameters
	params := []string{
		"pocketsphinx_continuous",
		"-alpha", "0.97",
		"-backtrace", "yes",
		"-beam", "1e-10000",
		"-bestpath", "no",
		"-cmn", "live",
		"-cmninit", "52.55,0.14,-3.23,14.29,-7.74,9.03,-7.17,-6.31,-0.13,1.09,5.23,-2.69,1.01",
		"-dict", "/home/dbarbera/Repositories/mySphinx/data/art_db.phone",
		"-dither", "no",
		"-doublebw", "no",
		"-featparams", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us/feat.params",
		//"-featparams", "/usr/local/share/pocketsphinx/model/en-us/en-us/feat.params",
		"-fsgusefiller", "no",
		"-frate", "125",
		"-fwdflat", "no",
		"-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
		//"-hmm", "/usr/local/share/pocketsphinx/model/en-us/en-us",
		"-infile", "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav",
		"-jsgf", "/home/dbarbera/Repositories/sphinx/data/kl_ay_m.jsgf",
		"-logfn", "/home/dbarbera/Repositories/mySphinx/data/output.log",
		"-lpbeam", "1e-10000",
		"-lponlybeam", "1e-10000",
		"-lw", "6",
		"-maxhmmpf", "-1",
		"-maxwpf", "-1",
		"-nfft", "256",
		"-nwpen", "1",
		"-pbeam", "1e-10000",
		"-pip", "1",
		"-pl_window", "0",
		"-remove_dc", "no",
		"-remove_noise", "yes",
		"-remove_silence", "no",
		"-topn", "4",
		"-vad_postspeech", "25",
		"-vad_prespeech", "5",
		"-vad_startspeech", "8",
		"-vad_threshold", "1",
		"-wbeam", "1e-10000",
		"-wip", "0.5",
		"-wlen", "0.016"}

	params2 := []string{
		"pocketsphinx_continuous",
		"-alpha", "0.97",
		"-backtrace", "yes",
		"-beam", "1e-10000",
		"-bestpath", "no",
		"-cmn", "live",
		"-cmninit", "43.46,-0.55,-4.37,11.73,-6.42,8.67,-8.58,-7.35,-0.16,2.92,6.63,0.05,4.06",
		"-dict", "/home/dbarbera/Repositories/mySphinx/data/art_db.phone",
		"-dither", "no",
		"-doublebw", "no",
		//"-featparams", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us/feat.params",
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
		"-fsgusefiller", "no",
		"-frate", "125",
		"-fwdflat", "no",
		//"-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
		"-hmm", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us",
		"-infile", "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav",
		"-jsgf", "/home/dbarbera/Data/climb/forced_align_5a02788e-3d91-446f-828a-428b7f2d8785_climb1_colin_fixed_trimmed.wav.jsgf",
		"-logfn", "/home/dbarbera/Data/climb/output/output.log",
		"-lpbeam", "1e-10000",
		"-lponlybeam", "1e-10000",
		"-lw", "6",
		"-maxhmmpf", "-1",
		"-maxwpf", "-1",
		"-nfft", "256",
		"-nwpen", "1",
		"-pbeam", "1e-10000",
		"-pip", "1",
		"-pl_window", "0",
		"-remove_dc", "no",
		"-remove_noise", "yes",
		"-remove_silence", "no",
		"-topn", "4",
		"-vad_postspeech", "25",
		"-vad_prespeech", "5",
		"-vad_startspeech", "8",
		"-vad_threshold", "1",
		"-wbeam", "1e-10000",
		"-wip", "0.5",
		"-wlen", "0.016"}

	params3 := []string{
		"pocketsphinx_continuous", "-nwpen", "1", "-backtrace", "yes", "-maxwpf", "-1", "-lw", "6",
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
		"-dict", "/home/dbarbera/Data/art_db.phone", "-fwdflat", "no", "-wlen", "0.016", "-frate", "125", "-wbeam", "1e-10000", "-remove_silence", "no", "-vad_postspeech", "25", "-doublebw", "no", "-vad_threshold", "1", "-fsgusefiller", "no",
		"-jsgf", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_7302731c-188f-4dfe-83ce-a5a004f1cab2/forced_align_450654b3-3a8f-4709-821f-9341848ccd86_climb1_colin_fixed_trimmed.wav.jsgf", "-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
		"-infile", "/home/dbarbera/Repositories/test_pronounce/audio_clips/climb1_colin_fixed_trimmed.wav",
		"-cmninit", "43.46,-0.55,-4.37,11.73,-6.42,8.67,-8.58,-7.35,-0.16,2.92,6.63,0.05,4.06", "-vad_prespeech", "5", "-dither", "no", "-topn", "4",
		"-remove_noise", "yes", "-remove_dc", "no", "-nfft", "256",
		"-logn", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_7302731c-188f-4dfe-83ce-a5a004f1cab2/3a039639-a28c-4a17-80d9-260eb04a5ef2.log",
		"-cmn", "-live", "-wip", "0.5",
	} //Just

	//result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

	jsgf_buffer, err := os.ReadFile(filename_jsgf)
	check(err)
	audio_buffer, err := os.ReadFile(filename_wav)
	check(err)

	//Checks:
	fmt.Printf("fjsgf %T, %d\n", jsgf_buffer, len(jsgf_buffer))
	fmt.Printf("fwav %T, %d\n", audio_buffer, len(audio_buffer))
	fmt.Printf("Params: %T, %d\n", params, len(params2))

	fmt.Printf("%s\n", filename_jsgf[:len(filename_jsgf)-5]+"__from_c.jsgf")
	fmt.Printf("%s\n", filename_wav[:len(filename_wav)-4]+"__from_c.jsgf")

	//xyz.Create_file(jsgf_buffer, filename_jsgf[:len(filename_jsgf)-5]+"__from_c.jsgf")
	//xyz.Create_file(jsgf_buffer, filename_wav[:len(filename_wav)-4]+"__from_c.wav")

	//xyz.Create_file_params_nofilename(params)
	//xyz.Check_string("Hello World, passing go strings to c strings, yeahhh again and again!!!")
	//xyz.Passing_bytes(audio_buffer)
	//xyz.Create_file_params(params, "./../data/params_file__from_c.txt")
	// xyz.Ps_call(jsgf_buffer, audio_buffer, params)

	// fmt.Printf("Results -> type: %T, length: %s\n", few_strings, len(few_strings))
	// fmt.Printf("Params: %T, %d\n", params, len(params))

	str := []string{"aaaabbbbccc"}
	fmt.Println(str)
	xyz.Modify_string(str)
	fmt.Println(str)

	//New Ps_call
	//xyz.Ps_call(jsgf_buffer, audio_buffer, params, result)

	// //fmt.Println(result)
	// var gostring = "hello from go!!!"
	// xyz.Check_string(gostring)

	// xyz.Create_file_params_nofilename(params)

	//xyz.Mock_ps_call()

	//New Ps_call
	// fmt.Println(result)
	// start := time.Now()
	// xyz.Ps_call(jsgf_buffer, audio_buffer, params, result)
	// elapsed := time.Since(start)
	// fmt.Printf(">>>> Timing for one call: %s\n", elapsed)
	// fmt.Println(result)

	// ncalls := 10
	// startn := time.Now()
	// for i := 0; i < ncalls; i++ {
	// 	//fmt.Println(result)
	// 	xyz.Ps_call(jsgf_buffer, audio_buffer, params, result)
	// 	//fmt.Println(result)
	// }
	// elapsedn := time.Since(startn)
	// fmt.Printf(">>>> Timing for %d calls: %s\n", ncalls, elapsedn)
	// fmt.Println(result)

	//r := []xyz.Utt{}
	starti := time.Now()
	var r = Ps(jsgf_buffer, audio_buffer, params2)
	elapsedi := time.Since(starti)
	fmt.Printf(">>>> Timing: %s\n", elapsedi)
	fmt.Println(r)

	fmt.Println(params2)
}
