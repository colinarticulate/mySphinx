/* File : example.i */
%module example

%{
extern int passing_bytes(char *bytes, int len);
%}

%typemap(gotype) (char *bytes, int len) "[]byte"

%typemap(in) (char *bytes, int len)
%{
  $1 = $input.array;
  $2 = $input.len;
%}

extern int passing_bytes(char *bytes, int len);

