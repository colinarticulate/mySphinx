package main

import (
	"fmt"
	"os"
	//"os"
	//"xyz"
	//"github.com/davidbarbera/articulate-pocketsphinx-go/xyz"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main_() {
	//Data:
	//Filenames
	filename_jsgf := "./../data/kl_ay_m.jsgf"
	filename_wav := "./../data/climb1_colin.wav"
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
		"-fsgusefiller", "no",
		"-frate", "125",
		"-fwdflat", "no",
		"-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
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

	few_strings := []string{"NULL", "null", "NULL", "null", "NULL", "null", "NULL", "null", "NULL",
		"null", "NULL", "null", "NULL", "null", "NULL", "null", "NULL", "null", "NULL", "null", "NULL", "null"}

	jsgf_buffer, err := os.ReadFile(filename_jsgf)
	check(err)
	audio_buffer, err := os.ReadFile(filename_wav)
	check(err)

	//Checks:
	fmt.Printf("fjsgf %T, %d\n", jsgf_buffer, len(jsgf_buffer))
	fmt.Printf("fwav %T, %d\n", audio_buffer, len(audio_buffer))
	fmt.Printf("Params: %T, %d\n", params, len(params))

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

	// xyz.Modify_strings(few_strings)

	n := len(few_strings)

	fmt.Printf("Results -> type: %T, length: %d\n", few_strings, len(few_strings))
	for i := 0; i < n; i++ {
		fmt.Println(few_strings[i])
	}

	// fmt.Println("this is a tab \t\t\t, and this a \n.")

	// test := []string{"sil kl ay m v b sil (-3641)*word,start,end*sil,3,90*(NULL),90,90*kl,91,109*(NULL),109,109*ay,110,147*(NULL),147,147*m,148,165*(NULL),165,165*v,166,172*b,173,176*sil,177,245**0000000000000000000000000000000000000"}

	// raw := strings.split(test[0], "**")

}
