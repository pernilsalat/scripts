#int Rutina (int v[20], int limite) {
#	int i=1, resultado=0;
#	while (i<limite) {
#		if (v[i]==5) funcion(&resultado)
#		++i
#	}
#	return resultado;
#}

  pushl %ebp
  movl %esp, %ebp
  pushl $0
  pushl %esi
  pushl %edi
  pushl %ebx
  movl 8(%ebp), %esi  # start v
  movl 12(%ebp), %edi
  shll $2, %edi
  addl %esi, %edi     # fin v
  leal -4(%ebp), %ebx # &result

while:
  cmpl %esi, %edi
  je fi_while
if:
  cmpl $5, (%esi)
  jne fi_if
  pushl %ebx
  call funcion
  addl $4, %esp
fi_if:
  addl $4, %esi
fi_while:
  popl %ebx
  popl %edi
  popl %esi
  movl -4(%ebp), %eax
  movl %ebp, %esp
  popl %ebp
  ret
