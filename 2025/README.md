# Idea rápida

Una recta que sólo tiene los valores (0,1,\dots,99), entonces todos los cálculos se hacen **mod (100)**.
Eso significa que dos enteros (x) e (y) son equivalentes si difieren por un múltiplo de (100):
[
x \equiv y \pmod{100} \quad\Longleftrightarrow\quad x-y \in 100\mathbb{Z}.
]

# Demostración informal

Queremos calcular (50-68) en la recta ({0,\dots,99}).

1. Calcula en (\mathbb{Z}): (50-68=-18).
2. Equivalente en el conjunto ({0,\dots,99}). Como (-18) no está en ese conjunto, sumamos (100) (un múltiplo del módulo) hasta obtener un número en ([0,99]):
   [
   -18 + 100 = 82.
   ]
3. Por tanto (50-68\equiv -18\equiv 82\pmod{100}). En la recta circular el resultado es (82).

# Fórmula general

Para enteros (a,b) y módulo (m=100), el resultado reducido al intervalo ({0,\dots,m-1}) es
[
r \equiv a-b \pmod{m},\qquad r\in{0,\dots,m-1}.
]
Una forma cerrada es
[
r = (a-b) - m\left\lfloor\frac{a-b}{m}\right\rfloor,
]
que produce el representante único (r) en ([0,m-1]). En la práctica, si (a-b\ge0) entonces (r=a-b); si (a-b<0) entonces (r=a-b+m) (o sumar (m) suficientes veces si la diferencia es menor que (-m)).

# Justificación de unicidad

Por el teorema de la división, para cualquier entero (x) existen únicos (q,r) con (x=mq+r) y (0\le r<m). Ese (r) es el representante único en ({0,\dots,m-1}) congruente con (x) módulo (m).

# Algoritmo

Para calcular en la recta de (0..99):

1. Calcula (d=a-b).
2. Si (0\le d\le99) esa es la respuesta.
3. Si (d<0), suma (100) hasta que el resultado caiga en ([0,99]): (d' = d + 100).
4. Si (d\ge100), resta (100) hasta reducirlo a ([0,99]).

e.g: (50-68=-18) → (-18+100=82).

