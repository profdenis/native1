# Chapitre 1 : Les bases du langage Go

## Section 1 : Le premier programme

Ce n'est pas original, mais le premier programme dans un nouveau langage de programmation est presque toujours le fameux
*Hello World !*.

````go
package main

import "fmt"

func main() {
	fmt.Println("Hello World !")
}
````

### Explications

1. `package main` : tous les fichiers `.go` doivent débuter par une déclaration de paquet (*package*), et lorsqu'on veut
   un *exécutable*, le paquet doit être `main`, et il doit y avoir une fonction appelée `main`.
2. `import "fmt"` : on doit importer le paquet `fmt` pour pouvoir utiliser la fonction `Println` pour écrire du texte
   sur la sortie standard (*standard output*, ou *stdout*).
3. `func main() {` : déclaration de la fonction `main`, qui sera le point de départ de notre programme exécutable.
4. `fmt.Println("Hello World !")` : le contenu de la fonction est évidemment très simple, on écrit simplement une chaîne
   de caractères sur le *stdout*
    1. `Print` --> imprime (ou écrit), `ln` --> ajoute un changement de ligne (un `\n`) à la fin
    2. les fonctions et les variables **publiques** commencent toujours avec une lettre majuscule en Go : si la fonction
       était définie avec une lettre minuscule au début, soit `println` à la place, la fonction serait **privée** et on
       ne pourrait pas l'appeler
    3. il existe plusieurs autres fonctions dans le paquet `fmt`, pour écrire sur le *stdout* ou ailleurs
5. `}` : les blocs de code sont délimités par des accolades (comme en Java), mais il n'y a pas de point-virgules `;` à
   la fin des lignes qui contiennent des énoncés qui ne nécessitent pas de définition de bloc

## Section 2 : les variables et les types de base

Référence : [A tour of Go](https://go.dev/tour/basics/11)

| catégorie                             | types                                                    |
|---------------------------------------|----------------------------------------------------------|
| logique                               | `bool`                                                   |
| caractères                            | `string`, `rune` (alias pour `int32`)                    |
| nombres entiers (signés)              | `int`, `int8`, `int16`, `int32`, `int64`                 |
| nombres entiers positifs (non signés) | `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` |
| données brutes                        | `byte` (alias for `uint8`)                               |
| nombres réels                         | `float32`, `float64`                                     |
| nombres complexes                     | `complex64`, `complex128`                                |

Les types `int`, `uint`, et `uintptr` ont une largeur de 32 bits sur un système 32-bit et une largeur de 64 bits sur un
système 64-bit. Quand vous avez besoin d'un nombre entier, vous devriez utiliser `int` à moins d'avoir une raison
spécifique d'utiliser une taille précise ou un nombre non-signé.

### Définition avec `var` et sans valeur initiale

Une déclaration de type `var` commence avec le mot-clé `var`, suivi par le nom de la variable ou une liste de noms de
variables séparées par des virgules, et suivi par le type de la ou les variable(s).

````go
package main

import "fmt"

func variables2a() {
	var i int
	var finished bool
	var name string
	var x, y float32

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
}
````

Les variables vont être automatiquement initialisées avec leur **valeur zéro** (ou leur valeur par défaut), qui dépend
du type de la variable. Pour les nombres, la valeur par défaut est `0`, pour les booléens la valeur est `false`, et pour
les chaînes de caractères, la chaîne vide `""` est utilisée comme valeur zéro.

### Définition avec `var` et valeur initiale

On peut aussi préciser une valeur pour initialiser une variable au moment de la déclaration.

````go
package main

import "fmt"

func variables2b() {
	var i int = 5
	var finished bool = true
	var name string = "Denis"
	var x, y float32 = 2.5, -1.4

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
}
````

Si une valeur initiale est spécifiée, alors il n'est pas nécessaire de spécifier le type. La variable prendra
automatiquement le type de la valeur utilisée. L'exemple suivant est équivalent à l'exemple précédent, à une différence
près : les variables `x` et `y` sont maintenant de type `float64`, et non `float32`. Le type pour les nombres réels
est `float64` par défaut.

````go
package main

import "fmt"

func variables2c() {
	var i = 5
	var finished = true
	var name = "Denis"
	var x, y = 2.5, -1.4

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
}
````

#### Constantes

On peut définir une constante en remplaçant `var` par `const`. Essayer de modifier la valeur d'une constante va créer
une erreur de compilation.

### Définition avec le raccourci `:=`

Il est possible d'utiliser un raccourci pour les définitions de variables avec valeurs initiales, à condition de
définir seulement une variable à la fois. L'exemple suivant est équivalent à l'exemple précédent, sauf que `x` et `y`
sont définies séparément.

````go
package main

import "fmt"

func variables2d() {
	i := 5
	finished := true
	name := "Denis"
	x := 2.5
	y := -1.4

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
}
````

## Section 3 : les différentes façons d'écrire des valeurs

En plus des fonctions `Print`, `Println` et `Printf` qui sont utilisées pour écrire sur le _stdout_, le paquet `fmt`
contient des fonctions pour écrire dans une chaîne de caractères (`Sprint`, `Sprintln`, `Sprintf`) ou dans un
fichier (`Fprint`, `Fprintln`, `Fprintf`).

- Les fonctions se terminant par *ln* ajoutent un changement de ligne (`\n`) à la fin de la chaîne de caractères à
  écrire.
- Les fonctions se terminant par un *f* utilisent une syntaxe spéciale pour **formatter** les chaînes de caractères.
- Les fonctions commençant avec un *S* écrivent dans une chaîne de caractères (*S* pour *string*)
- Les fonctions commençant avec un *F* écrivent dans un fichier (*F* pour *file*)
    - Les fonctions sur les fichiers seront présentées plus loin, dans le chapitre consacré à la gestion des fichiers.

### Écrire dans une chaîne de caractères

````go
package main

import "fmt"

func print1() {
	year := 2024
	output := fmt.Sprintln("année =", year)
	fmt.Println(output)
}
````

Dans cet exemple simplifié, il serait préférable d'utiliser `Println` directement sur les deux valeurs à
écrire (`fmt.Pprintln("année =", year)`), mais dans un exemple plus complexe, lorsqu'il est nécessaire d'écrire la même
chaîne de caractères à plusieurs reprises, possiblement à des endroits différents (_stdout_, fichier, ...), écrire dans
une chaîne en premier peut simplifier le code plus loin dans le programme, surtout si on a besoin de formatter les
valeurs de façon particulière.

### Écrire selon un format

Le même principe s'applique aux fonctions `Printf`, `Sprintf` et `Fprintf` : une chaîne de caractères décrivant le
format doit être donnée comme premier argument aux fonctions, suivie d'une ou plusieurs valeurs à formatter.

#### Nombres entiers

````go
package main

import "fmt"

func print2() {
	year := 2024
	month := 1
	day := 5
	fmt.Printf("année = %d\n", year)
	fmt.Printf("%d-%d-%d\n", year, month, day)
	fmt.Printf("%4d-%2d-%2d\n", year, month, day)
	fmt.Printf("%4d-%02d-%02d\n", year, month, day)
}
````

Sortie :

````
année = 2024
2024-1-5
````

La chaîne de caractères décrivant le format (_chaîne de format_, ou _format string_ en anglais) doit inclure au moins
une entrée débutant par le caractère `%`, suivi par la description du format. Dans l'exemple précédent, il y a 2
sorties.
Dans la première, la sortie sera identique à l'exemple précédent. Le `%d` va être remplacé par la valeur de la
variable `year`. Le `d` indique que la valeur sera un nombre entier en base 10 (_décimal_). On pourrait utiliser

- `%x` pour un nombre entier en base 16 (_hexadécimal_)
- `%o` pour un nombre entier en base 8 (_octal_)
- `%b` pour un nombre entier en base 2 (_binaire_)

Dans la deuxième sortie, il y a 3 formats à l'intérieur de la chaîne de format, et ils sont séparés par des tirets qui
ne font pas partie des formats, mais qui feront partie de la sortie. Le problème est que la sortie n'est pas très belle
parce que les valeurs pour le mois et la journée sont plus petites que 10, alors qu'habituellement, on réserve 2
emplacements (ou 2 colonnes) pour ces valeurs même si elles n'en n'ont pas besoin.

On peut remplacer le deuxième format par celui-ci : `"%4d-%2d-%2d\n"`. La sortie sera alors `2024- 1- 5`, ce qui n'est
pas vraiment mieux. On spécifie que l'on veut avoir 4 colonnes pour l'année, 2 colonnes pour le mois et aussi 2 colonnes
pour le jour, mais par défaut, les colonnes libres sont remplies avec des espaces.

Si on remplace le deuxième format par celui-ci : `"%4d-%02d-%02d"`, alors les colonnes libres seront remplies par des
zéros, ce qui correspond au format habituel `2024-01-05`.

Cet exemple montre comment formatter une date dans un format précis, mais il existe plusieurs autres formats possibles.
Formatter des dates est un besoin commun, qui doit supporter des formats différents selon les paramètres
de localisation de l'application. Le paquet `time` peut être utilisé pour, entre autres, formatter les dates et
les heures. Pour plus de détails :
[Date and time format in Go cheatsheet](https://gosamples.dev/date-time-format-cheatsheet/)

#### Formatter d'autres types

Référence : [Paquet `fmt`](https://pkg.go.dev/fmt)

Les formats les plus courants :

| Formats          | Détails                                     |
|------------------|---------------------------------------------|
| `%v`             | format par défaut                           |
| `%d`, `%x`, `%b` | nombres entiers en base 10, 16 et 2         |
| `%f`, `%F`       | nombres réels, avec un point, sans exposant |
| `%e`, `%e`       | nombres réels, notation scientifique        |
| `%s`             | chaînes de caractères                       |
| `%t`             | booléen                                     |
| `%T`             | type de la valeur                           |

Pour les nombres réels, on peut spécifier la précision voulue.

````go
package main

import "fmt"

func print3() {
	x := 1234.5678
	fmt.Printf("%f\n%.2f\n%.f\n", x, x, x)
}
````

Sortie :

````
1234.567800
1234.57
1235
````

## Section 4 : les fonctions

Une déclaration de fonction a la structure suivante en Go :

````go
func nomFonction(param1 type1, param2 type2, ...) typeRetour {}
````

Une déclaration commence toujours par `func`, suivi du nom de la fonction, qui doit être écrite en `camelCase` ou
en `PascalCase`. Comme pour les variables, si le nom de la fonction commence par une **minuscule**, la fonction sera 
**privée** et accessible seulement dans le paquet qui la contient, tandis que si le nom commence par une **majuscule**,
la fonction sera **publique**, donc accessible de l'extérieur du paquet. Donc, pour une fonction privée, on utilise le
_camelCase_, et pour une fonction publique, on utilise le _PascalCase_.

Si la fonction ne retourne rien, alors on doit omettre le type de retour. On ne doit **pas** utiliser `void` comme dans
d'autres langages.

````go
package main

import "fmt"

func allo(name string) {
	fmt.Printf("Allo %s !\n", name)
}
````

Sortie :

````
Allo Alice !
````

Voici une fonction qui fait un calcul simple et retourne le résultat :

````go
package main

import "fmt"

func square(x int) int {
	return x * x
}
````

Sortie avec l'appel `fmt.Println(square(4))` : `16`

Il est possible pour une fonction de retourner plus qu'une valeur. Les types de retours doivent être séparés par des
virgules, et des parenthèses doivent être utilisées autour des types de retour.

````go
func nomFonction(param1 type1, param2 type2, ...) (typeRetour1, typeRetour2, ...) {}
````

Voici une fonction qui retourne les 2 valeurs données en paramètre, mais en ordre inverse.

````go
package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}
````

Appel de la fonction :

````go
a, b := 2, 5
fmt.Println(swap(a, b))
````

Sortie : `5 2`

Voici une autre fonction pour échanger des valeurs, mais elle ne fonctionne pas :

````go
package main

import "fmt"

func swap2(x int, y int) {
	temp := x
	x = y
	y = temp
}
````

Appel de la fonction :

````go
a, b := 2, 5
fmt.Println(a, b)
swap2(a, b)
fmt.Println(a, b)
````

Sortie :

````
2 5
2 5
````

Plus de détails vont être donnés dans le chapitre suivant, mais la raison pourquoi cette fonction n'a aucun effet sur
les variables `a` et `b` est que la fonction travaille sur des copies des valeurs de `a` et `b`. Les variables  `x`
et `y` sont initialisés avec les valeurs de `a` et `b`, mais sont indépendantes de `a` et `b`.

La meilleure façon d'échanger les valeurs de deux variables en Go est celle-ci : `b, a = a, b`. `b` va obtenir la valeur
de `a`, et `a` va obtenir la valeur de `b`. Ça fonctionne parce que les changements de valeurs vont se faire en
parallel.

## Section 5 : les conditionnelles

### `if ... else`

La particularité des conditionnelles dans le langage Go est qu'on ne place pas la condition entre parenthèses, et que
les accolades sont obligatoires autour du code à exécuter dans la partie `if` et dans la partie `else`, même s'il y a
seulement une ligne de code à exécuter. De plus, l'accolade d'ouverture `{` doit être sur la même ligne que le `if` ou
le `else`.

Utilisez les opérateurs `&&` pour la conjonction (et), `||` pour la disjonction (ou) et `!` pour la négation. Ajouter
des parenthèses au besoin pour s'assurer que les opérateurs sont appliqués dans le bon ordre.

````go
package main

import "fmt"

func conditional1(x int) {
	if x > 0 {
		fmt.Printf("x = %d est plus grand que 0\n", x)
	} else {
		fmt.Printf("x = %d est plus petit ou égal à 0\n", x)
	}
}
````

Sortie pour des appels avec les valeurs 5, 0, et -5 :

````
x = 5 est plus grand que 0
x = 0 est plus petit ou égal à 0
x = -5 est plus petit ou égal à 0
````

Il n'y a pas de syntaxe particulière pour les séquences de `if ... else`, il faut simplement démarrer un autre
`if ... else` tout de suite après le premier `else`.

````go
package main

import "fmt"

func conditional2(x int) {
	if x > 0 {
		fmt.Printf("x = %d est plus grand que 0\n", x)
	} else if x == 0 {
		fmt.Println("x est égal à 0")
	} else {
		fmt.Printf("x = %d est plus petit que 0\n", x)
	}
}
````

Sortie pour des appels avec les valeurs 5, 0, et -5 :

````
x = 5 est plus grand que 0
x est égal à 0
x = -5 est plus petit que 0
````

### `switch`

Un _switch_ (_aiguillage_) est une bonne façon de simplifier la syntaxe quand on a besoin d'une série de `if ... else`.
Il existe 2 principales formes de `switch` en Go : un `switch` sur une variable avec des cas sur différentes valeurs
possibles de la variable, et un `switch` sans variable avec des cas sur des expressions booléennes.

#### `switch` sans variable

````go
package main

import "fmt"

func conditional3(x int) {
	switch {
	case x > 0:
		fmt.Printf("x = %d est plus grand que 0\n", x)
	case x == 0:
		fmt.Println("x est égal à 0")
	case x < 0:
		fmt.Printf("x = %d est plus petit que 0\n", x)
	}
}
````

Sortie pour des appels avec les valeurs 5, 0, et -5 : même que l'exemple précédent.

Cet exemple correspond presque exactement à l'exemple précédent écrit avec des `if ... else`, sauf pour le dernier cas
qui pourrait être remplacé par le cas spécial `default` qui correspond au dernier `else` d'une série de `if ... else`.
On obtiendra la même sortie avec ce changement.

Dans ce cas-ci, remplacer `case x < 0` par `default` est approprié parce que la condition `x < 0` couvre toutes les
possibilités restantes. Notez qu'il n'est pas nécessaire d'utiliser des `break` pour terminer chaque cas.

````go
package main

import "fmt"

func conditional4(x int) {
	switch {
	case x > 0:
		fmt.Printf("x = %d est plus grand que 0\n", x)
	case x == 0:
		fmt.Println("x est égal à 0")
	default:
		fmt.Printf("x = %d est plus petit que 0\n", x)
	}
}
````

#### `switch` avec variable

````go
package main

import "fmt"

func conditional5(name string) {
	switch name {
	case "Denis":
		fmt.Println("Bonjour Prof. Denis !")
	case "":
		fmt.Println("Bonjour Inconnu !")
	default:
		fmt.Printf("Allo %s !\n", name)
	}
}
````

Sortie pour des appels avec les valeurs `"Denis"`, `""`, et `"Alice"` :

````
Bonjour Prof. Denis !
Bonjour Inconnu !
Allo Alice !
````

Lorsqu'on _switch_ sur une variable, les cas doivent être des valeurs spécifiques de la variable. Si on veut grouper
plusieurs valeurs dans un même cas, on peut les séparer par des virgules, comme dans l'exemple suivant.

````go
package main

import "fmt"

func conditional6(name string) {
	switch name {
	case "Denis", "Benoit":
		fmt.Printf("Bonjour Prof. %s !\n", name)
	case "":
		fmt.Println("Bonjour Inconnu !")
	default:
		fmt.Printf("Allo %s \n!", name)
	}
}
}
````

Sortie pour des appels avec les valeurs `"Denis"` et `"Benoit"` :

````
Bonjour Prof. Denis !
Bonjour Prof. Benoit !
````

## Section 6 : les boucles

### Boucle avec un compteur

À la base, les boucles _pour_ en Go ressemblent aux boucles _pour_ dans d'autres langages comme C et Java, avec quelques
différences :

- il n'y a pas de parenthèses `()` autour des 3 composantes principales de la
  boucle _pour_: `initialisation ; condition ; incrémentation`
- le compteur de la boucle est normalement défini avec la notation abrégée utilisant le symbole `:=`
- les 3 composantes principales sont optionnelles

````go
package main

import "fmt"

func loop1() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

````

Sortie :

````
0 1 2 3 4
````

### Les boucles de type _while_

Techniquement, il n'y a pas de boucles _while_ en Go. Le seul type de boucle utilise _for_, comme démontré dans
l'exemple précédent. Puisque les parties _initialisation_ et _incrémentation_ sont optionnelles, on peut écrire une
boucle uniquement avec la condition.

````go
package main

import "fmt"

func loop2() {
	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
}

````

La sortie est la même que l'exemple précédent.

### Boucle sans condition

Voici une **très mauvaise** manière d'écrire une boucle qui produit le même résultat, qui est techniquement légale, elle
obéit à la syntaxe du langage, mais qui est beaucoup plus difficile à lire, et qui peut causer des problèmes de
maintenance de code plus tard.

````go
package main

import "fmt"

func loop3() {
	i := 0
	for {
		fmt.Print(i, " ")
		i++
		if i >= 5 {
			break
		}
	}
	fmt.Println()
}

````

À première vue, la boucle ressemble à une boucle infinie parce qu'il n'y a pas de condition après le `for`, ce qui est
légal en Go. C'est la même chose que s'il y avait une condition qui est toujours vraie, comme `for true` (ou comme
`while (true)` en Java). Mais cette boucle n'est pas vraiment infinie parce qu'il y a une condition à l'intérieur qui
devient vraie à un moment donné et qui fait un `break` pour sortir de la boucle.

En général, il faut éviter les boucles infinies ou qui en apparence ont l'air infinie parce qu'elles sont plus
difficiles à lire et à déboguer. Les vraies boucles infinies sont utiles dans certaines situations très particulières,
comme l'implémentation d'un serveur qui doit attendre un nombre indéfini de connexions ou de requêtes, mais ces
situations particulières sont rares.

### Boucles sur les tableaux

Voici comment définir un tableau de nombres entiers avec valeurs initiales, et comment faire une boucle sur le tableau
en utilisant les index, ou positions, des éléments dans le tableau, pour écrire tous les éléments du tableau sur le
_stdout_ avec un format particulier. On utilise la fonction `len` pour obtenir la longueur d'un tableau.

On peut également donner un tableau directement à la fonction `Println` pour en écrire le contenu.

````go
package main

import "fmt"

func loop4() {
	numbers := [7]int{2, 5, 3, 4, 7, 1, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}
	fmt.Println(numbers)
}
````

Sortie :

````
numbers[0] = 2
numbers[1] = 5
numbers[2] = 3
numbers[3] = 4
numbers[4] = 7
numbers[5] = 1
numbers[6] = 7
[2 5 3 4 7 1 7]
````

Si on n'initialise pas le tableau immédiatement avec des valeurs spécifiées entre accolades immédiatement après le
type `[7]int`, alors le tableau sera initialisé avec les valeurs zéro du type, dans ce cas-ci `0` à cause du type `int`.
Il est possible de remplacer la longueur `7` par `...` pour que le compilateur calcule la longueur automatiquement. Il
est également possible de ne rien spécifier entre les `[]`, mais dans ce cas, on n'obtiendra pas un tableau, mais une
_slice_ (tranche). Les tableaux et les tranches se ressemblent, mais il y a des différences en termes de leurs
propriétés et de la gestion de la mémoire. Plus de détails viendront à ce sujet dans le prochain chapitre.

Il existe une meilleure façon de parcourir tous les éléments d'un tableau en Go : en utilisant `range`.

````go
package main

import "fmt"

func loop5() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	for i, number := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, number)
	}
	fmt.Println(numbers)
}
````

La sortie est la même que l'exemple précédent. À chaque tour de boucle, `i` a l'index courant, et `number` a la valeur
de l'élément courant du tableau. Les avantages de ce type de boucle est qu'on n'a pas besoin de gérer l'index
manuellement, il est incrémenté automatiquement, et la fin de boucle est vérifiée aussi automatiquement. Les
désavantages de ce type de boucle sont qu'on ne pas modifier le contenu du tableau dans la boucle, et on ne peut pas
parcourir les éléments du tableau dans un ordre différent (par exemple, de la fin vers le début).

Si on n'a pas besoin des index, mais seulement des valeurs des éléments dans le tableau, on doit alors faire la boucle
de la façon suivante.

````go
package main

import "fmt"

func loop6() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}
````

Sortie :

````go
2 5 3 4 7 1 7 
````

Dans le langage Go, toutes les variables déclarées doivent être utilisées au moins une fois. Puisque dans cet exemple,
on ne veut pas utiliser l'index des éléments du tableau, on utilise la barre de soulignement `_` à la place d'un nom de
variable. Si on écrivait `i` à la place de `_` dans cet exemple, on obtiendrait l'erreur `i declared and not used`.

## Références

1. [Go cheatsheet](https://devhints.io/go)
2. [A tour of Go](https://go.dev/tour/list)
3. [Paquet `fmt`](https://pkg.go.dev/fmt)



