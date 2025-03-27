# import random

# l1 = ["fsd", "sdfs", "vbcxvxn", "ytghj", "cvxccv"]
# l2 = ["fsd", "5645", "vbn", "xcvxc", "cvcv"]
# l3 = ["fsd", "5wer645", "uuuu", "ytghj", "cvcbcvbwev"]
# l4 = [1, 2, 3, 4, 5]
# l5 = [6, 7, 8, 9, 0]
# l6 = [11, 12, 13, 14, 15]
# l7 = [1.1, 1.2, 1.3, 1.4, 1.5]
# l8 = [2.1, 2.2, 2.3, 2.4, 2.5]
# l9 = [3.1, 3.2, 3.3, 3.4, 3.5]

# lists_combined = [l1, l2, l3, l4, l5, l6, l7, l8, l9]
# [print(mylist) for mylist in lists_combined]
# [print(mylist[random.randint(0, 4)]) for mylist in lists_combined]


namens_liste = [
    "Janine",
    "Alex",
    ["Markus", "Sanjar", "Obito", "Lia", "Rosa"],
    "Serhat",
    "Anja",
    "Marcel",
]

telefon_liste = [
    15,
    54,
    78,
    [
        16,
        18,
        25,
        73,
        19,
    ],
    32,
    17,
]

nummern_liste = [
    12.3,
    34.6,
    75.3,
    25.65,
    [
        15.5,
        13.32,
        14.0,
        11.33,
        5.05,
    ],
    30.00,
    111,
]
# Gebe dir nacheinander folgende Werte aus folgender Liste aus:

# 1. Aus Liste namens_liste den Namen Alex
print(namens_liste[1])
# 2. aus der Liste telefon_liste die Zahl 32
print(telefon_liste[-2])
# 3. Aus der Liste nummern_liste die Zahl 111
print(nummern_liste[-1])
# 4. Aus der Liste namens_liste den Namen Sanjar
print(namens_liste[2][1])
# 5. Aus der telefon_liste die Zahl 25
print(telefon_liste[3][2])
# 6. Aus der Liste nummern_liste die Zahl 13.32
print(nummern_liste[4][1])
# 7. Aus der Liste namens_liste den Namen Serhat
print(namens_liste[3])
# 8. Aus der telefon_liste die Zahl 19
print(telefon_liste[3][-1])
# 9. Aus der Liste nummern_liste die Zahl 5.05
print(nummern_liste[4][-1])
# 10. Aus der Liste namens_liste den Namen Anja
print(namens_liste[-2])
# 11. Aus der telefon_liste die Zahl 16
print(telefon_liste[3][0])
# 12. Aus der Liste nummern_liste die Zahl 30.00
print(f"{nummern_liste[-2]:.2f}")
print(
    [
        1,
        2,
        3,
        4,
        5,
        6,
        7,
        8,
    ][2::-1]
)


liste = [
    10,
    20,
    30,
    40,
    50,
    60,
    70,
    80,
    90,
    100,
]
print("aufgabe 1", liste[1:5])
print("aufgabe 2", liste[4:6])
print("aufgabe 3", liste[2:9])
print("aufgabe 4", liste[6:8])
print("aufgabe 5", liste[1:])
print("aufgabe 5", liste[:7])
print("aufgabe 6", liste[2:-3])
print("aufgabe 7", liste[4:-1])
print("aufgabe 8", liste[-6:-2])
print("aufgabe 9", liste[4:-3])
print("aufgabe 10", liste[1:-1:2])
print("aufgabe 11", liste[-2:-7:-3])
print("aufgabe 12", liste[-2:-9:-4])
print("aufgabe 13", liste[-1:-6:-1])
print("aufgabe 14", liste[1:7:3])
print("aufgabe 15", liste[-5:])
print("aufgabe 16", liste[:4])
print("aufgabe 17", liste[3:])
print("aufgabe 18", liste[::3])


def swapFirstLast(liste):
    first = liste[:-1]
    last = [liste.pop()]
    return last + first


print(swapFirstLast(liste))
