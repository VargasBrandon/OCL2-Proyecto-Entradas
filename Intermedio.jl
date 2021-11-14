x = 1::Int64;
y = 1::Int64;
println("---------------------------------");
println("Tablas de multiplicar con While");
println("---------------------------------");
while (x <= 10)
    while (y <= 10)
        print(x);
        print("x");
        print(y);
        print("=");
        println(x * y);
        global y = y + 1;
    end;
    println("-----------------------------");
    global x = x + 1;
    global y = 1;
end;

println("---------------------------------");
println("  Tablas de multiplicar con For");
println("---------------------------------");

for i in 1:10
    for j in 1:10
        print(i);
        print("x");
        print(j);
        print("=");
        println(i * j);
    end;
    println("--------------------------");
end;

iteraciones = 10::Int64;
temporal = 0::Int64;

while (temporal <= iteraciones)
    numero = temporal::Int64;
    if numero <= 0
        print("Factorial de ");
        print(temporal);
        println(" = 0");
        global temporal = temporal + 1;
        continue;
    end;
    factorial = 1::Int64;
    while (numero > 1)
        factorial = factorial * numero;
        numero = numero - 1;
    end;
    print("Factorial de ");
    print(temporal);
    print(" = ");
    println(factorial);
    temporal = temporal + 1;
end;


for i in 0:9

    output = "";
    for j in 0:(10 - i)
        output = output * " ";
    end;

    for k in 0:i 
        output = output * "* ";
    end;
    println(output);

end;