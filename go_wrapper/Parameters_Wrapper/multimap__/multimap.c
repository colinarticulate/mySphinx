/* File : example.c */
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>


int ps_call_params_only(int argc, char *argv[]){

  printf("%d parameters\n\n", argc);
  
  for( int i=0; i<argc; i++){
    printf("%d\t%s\n",i,argv[i]);
  }
return 0;

};