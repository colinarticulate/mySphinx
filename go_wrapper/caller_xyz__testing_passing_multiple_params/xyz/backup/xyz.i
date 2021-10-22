/* File : example.i */
//Swig 4.0 interface
%module xyz

// %{
// //extern int ps_call(char *bufferjsgf, int lenjsgf, char *bufferwav, int lenwav, char *argv[], int argc);
// //void create_file(char *buffer, int len, const char *filename);
// void create_file_params(char *argv[], int argc, const char *filename);
// %}

// // // %typemap(gotype) (char *bufferjsgf, int lenjsgf) "[]byte"
// // // %typemap(gotype) (char *bufferwav, int lenwav) "[]byte"
// // %typemap(gotype) (int argc, char *argv[]) "[]string"

// // // %typemap(in) (char *bufferjsgf, int lenjsgf)
// // // %{
// // //   $1 = $input.array;
// // //   $2 = $input.len;
// // // %}

// // // %typemap(in) (char *bufferwav, int lenwav)
// // // %{
// // //   $1 = $input.array;
// // //   $2 = $input.len;
// // // %}

// // %typemap(in) (char *argv[], int argc)
// // %{
// //     int i;
// //     _gostring_* a;

// //     $2 = $input.len;
// //     a = (_gostring_*) $input.array;
// //     $1 = (char **) malloc (($2 ) * sizeof (char *));

// //     for ( i = 0; i < $1; i++) {
// //       _gostring_ *ps = &a[i];
// //       $1[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
// //       memcpy($1[i],(char*) ps->p, (int)(ps->n) );
// //       $1[i][ps->n]='\0';
// //     }

// // %}

// // // //%typemap(out) (char *bufferjsgf, int lenjsgf, char *bufferwav, int lenwav, char *argv[], int argc) int
// // // //%typemap(argout) ( char *argv[], int argc) "" /* override char *[] default */

// // %typemap(freearg) (char *argv[], int argc)
// // %{
// //   for (i = 0; i < $2; i++) {
// //     free($1[i]);
// //   }
// //   free($1);
// // %}

// // // extern int ps_call(char *bufferjsgf, int lenjsgf, char *bufferwav, int lenwav, char *argv[], int argc);


// // %typemap(gotype) (char *buffer, int len) "[]byte"
// // %typemap(in) (char *buffer, int len)
// // %{
// //   $1 = $input.array;
// //   $2 = $input.len;
// // %}

// // extern void create_file(char *buffer, int len, const char *filename);

// %typemap(gotype) ( char *argv[], int argc) "[]string"

// %typemap(in) (char *argv[], int argc)
// %{
//     int i;
//     _gostring_* a;

//     $2 = $input.len;
//     a = (_gostring_*) $input.array;
//     $1 = (char **) malloc (($2 ) * sizeof (char *));
//     for ( i = 0; i < $2; i++) {
//       _gostring_ *ps = &a[i];
//       $1[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
//       memcpy($1[i],(char*) ps->p, (int)(ps->n) );
//       $1[i][ps->n]='\0';
//     }

// %}
// %typemap(argout) ( char *argv[], int argc) "" /* override char *[] default */

// %typemap(freearg) (char *argv[], int argc)
// %{
//   for (i = 0; i < $2; i++) {
//     free($1[i]);
//   }
//   free($1);
// %}

// void create_file_params( char *argv[], int argc, const char *filename);

// %{
// extern int create_file_params(int argc, char *argv[], char *filename);
// %}

// %typemap(gotype) (int argc, char *argv[]) "[]string"

// %typemap(in) (int argc, char *argv[])
// %{
//     int i;
//     _gostring_* a;

//     $1 = $input.len;
//     a = (_gostring_*) $input.array;
//     $2 = (char **) malloc (($1 ) * sizeof (char *));
//     for ( i = 0; i < $1; i++) {
//       _gostring_ *ps = &a[i];
//       $2[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
//       memcpy($2[i],(char*) ps->p, (int)(ps->n) );
//       //$2[i] = (char *) ps->p;
//       $2[i][ps->n]='\0';
//       // _gostring_ *ps = &a[i];
//       // $2[i] = (char *) ps->p;
//     }
//     //$2[i] = NULL;
//     //$2[i] = '\0';
  
// %}


// %typemap(argout) (int argc, char *argv[]) "" /* override char *[] default */

// %typemap(freearg) (int argc, char *argv[])
// %{
//   for (i = 0; i < $1; i++) {
//     free($2[i]);
//   }
//   free($2);
// %}

// %typemap(gotype) (char *filename) "string"

// %typemap(in) (char *filename) {
//   _gostring_ *ps = (_gostring_*) $input.p;
//   $1 = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
//   memcpy($1,(char*) ps->p, (int)(ps->n));
//   $1[ps->n]='\0';
// }

// %typemap(freearg) (char *filename) {
//   free($1);
// }

// extern int create_file_params(int argc, char *argv[], char *filename);


// //----------------------------------------
// //This works for check_string:
// %{
//  extern int check_string(char *c_string);
// %}

// %typemap(gotype) (char *c_string) "string"

// %typemap(in) (char *c_string) {
//   _gostring_ *ps = (_gostring_*) &$input;
//   $1 = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
//   memcpy($1,(char*) ps->p, (int)(ps->n));
//   $1[ps->n]='\0';
// }

// %typemap(freearg) (char *c_string) {
//   free($1);
// }

// extern int check_string(char *c_string);
// //----------------------------------------------------

%{
extern int create_file_params_nofilename(int argc, char *argv[]);
extern int check_string(char *str);
extern int create_file_params(int argc, char *argv[],  char *filename);
%}

//For argument filename:
%typemap(gotype) (char *str) "string"

%typemap(in) (char *str) {
  _gostring_ *ps = (_gostring_*) &$input;
  $1 = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
  memcpy($1, (char*)(ps->p), (int)(ps->n));
  $1[ps->n]='\0';
}

%typemap(freearg) (char *str) {
  free($1);
}


//For arguments argv and argc:
%typemap(gotype) (int argc, char *argv[]) "[]string"

%typemap(in) (int argc, char *argv[])
%{
    int i;
    _gostring_* a;

    $1 = $input.len;
    a = (_gostring_*) $input.array;
    $2 = (char **) malloc (($1 ) * sizeof (char *));
    for ( i = 0; i < $1; i++) {
      _gostring_ *ps = &a[i];
      $2[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
      memcpy($2[i],(char*) ps->p, (int)(ps->n) );
      //$2[i] = (char *) ps->p;
      $2[i][ps->n]='\0';
      // _gostring_ *ps = &a[i];
      // $2[i] = (char *) ps->p;
    }
    //$2[i] = NULL;
    //$2[i] = '\0';
  
%}

%typemap(argout) (int argc, char *argv[]) "" /* override char *[] default */

%typemap(freearg) (int argc, char *argv[])
%{
  for (i = 0; i < $1; i++) {
    free($2[i]);
  }
  free($2);
%}
extern int create_file_params_nofilename(int argc, char *argv[]);
extern int check_string(char *str);


//%typemap(gotype) (int argc, char *argv[]) "[]string", (char *str) "string"

%typemap(in) (int argc, char *argv[], char *str)
%{
    int i;
    _gostring_* a;

    $1 = $input[0].len;
    a = (_gostring_*) $input[0].array;
    $2 = (char **) malloc (($1 ) * sizeof (char *));
    for ( i = 0; i < $1; i++) {
      _gostring_ *ps = &a[i];
      $2[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
      memcpy($2[i],(char*) ps->p, (int)(ps->n) );
      //$2[i] = (char *) ps->p;
      $2[i][ps->n]='\0';
      // _gostring_ *ps = &a[i];
      // $2[i] = (char *) ps->p;
    }
    //$2[i] = NULL;
    //$2[i] = '\0';
    _gostring_ *ps = (_gostring_*) &$input[1];
    $1 = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
    memcpy($1, (char*)(ps->p), (int)(ps->n));
    $1[ps->n]='\0';
  
%}

%typemap(freearg) (int argc, char *argv[], char *str)
%{
  for (i = 0; i < $1; i++) {
    free($2[i]);
  }
  free($2);
  free($3);
%}


extern int create_file_params(int argc, char *argv[], char *str);


// //Exploring:
// %{
// extern int create_file_params(char *argv[], int argc, char *filename);
// %}

// //For argument filename:
// %typemap(gotype) (int argc, char *argv[]) "[]string"
// %typemap(gotype) (char *filename) "string"

// //%typemap(gotype) (int argc, char *argv[], char *filename) {"[]string", "string"}

// // %typemap(in) (char *filename) {
// //   _gostring_ *ps = (_gostring_*) &$input;
// //   $1 = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
// //   memcpy($1, (char*)(ps->p), (int)(ps->n));
// //   $1[ps->n]='\0';
// // }
// %typemap(in) (int argc, char *argv[], char* filename)
// %{
//   int i;
//   _gostring_* a;

//   $1 = $input1.len;
//   a = (_gostring_*) $input1.array;
//   $2 = (char **) malloc (($1 ) * sizeof (char *));
//   for ( i = 0; i < $1; i++) {
//     _gostring_ *ps = &a[i];
//     $2[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
//     memcpy($2[i],(char*) ps->p, (int)(ps->n) );
//     //$2[i] = (char *) ps->p;
//     $2[i][ps->n]='\0';
//     // _gostring_ *ps = &a[i];
//     // $2[i] = (char *) ps->p;
//   }
//   //$2[i] = NULL;
//   //$2[i] = '\0';

//   _gostring_ *ps_ = (_gostring_*) &$input2;
//   $3 = (char*)malloc( ((int)(ps_->n) + 1)*sizeof(char));
//   memcpy($3, (char*)(ps_->p), (int)(ps_->n));
//   $3[ps_->n]='\0';
 
// %}

// // %typemap(argout) (int argc, char *argv[]) "" /* override char *[] default */

// // %typemap(freearg) (char *filename) {
// //   free($1);
// // }
// // %typemap(freearg) (int argc, char *argv[])
// // %{
// //   for (i = 0; i < $1; i++) {
// //     free($2[i]);
// //   }
// //   free($2);
// // %}
// %typemap(freearg) (int argc, char *argv[], char *filename)
// %{
//   for (i = 0; i < $1; i++) {
//     free($2[i]);
//   }
//   free($2);
//   free($3)
// %}

// extern int create_file_params(char *argv[], int argc, char *filename);



//Solution might be here:
//https://stackoverflow.com/questions/65420683/swig-passing-multiple-arrays-from-python-to-c