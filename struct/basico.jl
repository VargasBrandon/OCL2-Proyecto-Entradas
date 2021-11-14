struct Personaje
  nombre::Int64;
  edad::Int64;
  descripcion::String;
end;

# Struct Mutable
mutable struct Carro
  placa::String;
  color::String;
  tipo::String;
end;

# Construcción Struct
p1 = Personaje("Fer", 18, "No hace nada");
p2 = Personaje("Fer", 18, "Maneja un carro");
c1 = Carro("090PLO", "gris", "mecanico");
c2 = Carro("P0S921", "verde", "automatico");

# Asignación Atributos
p1.edad = 10;
p2.edad = 20;
c1.color = "cafe";
c2.color = "rojo";

# Acceso Atributo
println(p1.edad);
println(c1.color);