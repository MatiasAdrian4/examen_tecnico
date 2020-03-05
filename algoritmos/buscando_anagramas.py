
def check_anagrama(a, b, pos):
    if pos + len(b) > len(a):
        return False
    sub_a = ""
    for i in range(len(b)):
        sub_a += a[pos+i]
    if ''.join(sorted(sub_a)) == b:
        print("Anagrama: ", sub_a)
        return True
    else:
        return False

def cantidad_anagramas(a, b):
    a = a.upper()
    b = ''.join(sorted(b.upper()))
    resultado = 0
    for i in range(len(a)):
        if a[i] in b:
            if check_anagrama(a, b, i):
                resultado = resultado + 1
    return resultado 


a = input("Ingrese la cadena principal: ")
b = input("Ingrese la subcadena a buscar: ")

# Ejemplo en el examen:
#a = "Hola, que buena ola laomir"
#b = "OAL"

print("Cantidad de Anagramas: ", cantidad_anagramas(a, b))