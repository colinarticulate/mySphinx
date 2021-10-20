/* File : multimap.i */
%module multimap

%{
extern int pscall(int argc, char *argv[]);
%}

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

extern int pscall(int argc, char *argv[]);




