/* File : example.i */
//Swig 4.0 interface
%module xyz


%{
extern int create_file_params_nofilename(int argc, char *argv[]);
extern int check_string(char *str);
extern int passing_bytes(char *bytes, int len);
%}

%typemap(gotype) (char *bytes, int len) "[]byte"

%typemap(in) (char *bytes, int len)
%{
  $1 = $input.array;
  $2 = $input.len;
%}

extern int passing_bytes(char *bytes, int len);

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



//Solution might be here:
//https://stackoverflow.com/questions/65420683/swig-passing-multiple-arrays-from-python-to-c
//That might work for Python but not for Go