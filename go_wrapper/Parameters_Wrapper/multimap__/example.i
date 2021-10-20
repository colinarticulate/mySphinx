/* File : example.i */
%module example

%{
extern int ps_call_params_only(int argc, char *argv[]);
%}

%typemap(gotype) (int argc, char *argv[]) "[]string"

%typemap(in) (int argc, char *argv[])
%{
  {
    int i;
    _gostring_* a;

    $1 = $input.len;
    a = (_gostring_*) $input.array;
    $2 = (char **) malloc (($1 + 1) * sizeof (char *));
    for (i = 0; i < $1; i++) {
      _gostring_ *ps = &a[i];
      $2[i] = (char *) ps->p;
    }
    $2[i] = NULL;
  }
%}

%typemap(argout) (int argc, char *argv[]) "" /* override char *[] default */

%typemap(freearg) (int argc, char *argv[])
%{
  free($2);
%}

extern int ps_call_params_only(int argc, char *argv[]);




