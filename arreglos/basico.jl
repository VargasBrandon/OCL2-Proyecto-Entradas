

arreglo1 = [1,2,3,4,5,6];
arreglo2 = [7,8,9,10,11,12,arreglo1];
arreglo3 = [13,14,15,16,16,18,arreglo2];


println(arreglo1[0]);
println(arreglo2[6][1]);
println(arreglo3[6][6][2]);

arreglo1[0]=4;
arreglo2[6][1]=5;
arreglo3[6][6][2]=6;

println(arreglo1[0]);
println(arreglo2[6][1]);
println(arreglo3[6][6][2]);

