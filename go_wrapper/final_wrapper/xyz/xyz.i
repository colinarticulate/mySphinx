/* File : example.i */
%module xyz

%{
extern int ps_call(char *bufferjsgf, int lenjsgf, char *bufferwav, int lenwav, char *argv[], int argc,);
%}

%typemap(gotype) (char *bufferjsgf, int lenjsgf) "[]byte"
%typemap(gotype) (char *bufferwav, int lenwav) "[]byte"
%typemap(gotype) (int argc, char *argv[]) "[]string"

%typemap(in) (char *bufferjsgf, int lenjsgf)
%{
  $1 = $input.array;
  $2 = $input.len;
%}
%typemap(in) (char *bufferwav, int lenwav)
%{
  $1 = $input.array;
  $2 = $input.len;
%}
%typemap(in) (char *argv[], int argc)
%{
    int i;
    _gostring_* a;

    $2 = $input.len;
    a = (_gostring_*) $input.array;
    $1 = (char **) malloc (($2 ) * sizeof (char *));

    for ( i = 0; i < $1; i++) {
      _gostring_ *ps = &a[i];
      $1[i] = (char*)malloc( ((int)(ps->n) + 1)*sizeof(char));
      memcpy($1[i],(char*) ps->p, (int)(ps->n) );
      $1[i][ps->n]='\0';
    }

%}

extern int ps_call(char *bufferjsgf, int lenjsgf, char *bufferwav, int lenwav, char *argv[], int argc,);

