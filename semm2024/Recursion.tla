---- MODULE recursion ----
EXTENDS Sequences, Naturals

VARIABLES element, list, cursor

Init == 
  /\ element = 0
  /\ list = <<>>
  /\ cursor = 0

AddElement(e) ==
    /\ element' = e
    /\ list' = Append(list, e)
    /\ cursor' = cursor

NextElement ==
    /\ element' = list[cursor]
    /\ list' = list
    /\ cursor' = cursor + 1

