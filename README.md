# go-concurrency-mutex

## Mutex 
Mechanism that allow us to prevent concurrent processes from entering a critical section or  that's already been executed  by another process.

![image](https://user-images.githubusercontent.com/32901911/110522043-3e8e8a00-80ef-11eb-86da-548c640671e1.png)


## Description of the problem

![image](https://user-images.githubusercontent.com/32901911/110520805-b52a8800-80ed-11eb-8b88-3ec25bf18821.png)

:blue_book: Imagine we have Gopher with a bank balance of a $1000 to start. then Gopher tries to add $ 500 to this account
Now one go routine would see this transaction, read the value of a $1000 and proceed to try and add $500 to that account.
That balance however at the exact same moment a charge of $700 as applied to pay his mortgage (hipoteca) and the second proccess reads the inital value of the $700 before the first is able to add the additional deposit of $500 and it then proceeds to substract $700 from a $1000.
Now the customer then checks his bank balance the next day, and notices that he is down $300. As the second process was unaware of the first deposit and overrruled the value upon completion of that second goroutine. So now that we know what the problem is let'see how we can fix it using a mutex 

:blue_book: Tenemos a Gopher que tiene en su cuenta, un balance de $1000. Quiere depositar $500, entonces una rutina va a leer el balance inicial y hacer el depósito de dicho monto.
Sin embargo, al mismo tiempo, se le aplica una extracción de $700 para pagar su hipoteca. Este segundo proceso, lo que hace es leer su depósito inicial $1000 y retirar los $700 respectivos.
Al otro día, Gopher chequea su cuenta y nota que tiene un balance de solamente $300 a favor.


![image](https://user-images.githubusercontent.com/32901911/110521358-66312280-80ee-11eb-94df-7c4790a57ba1.png)

Con Mutex, podemos solucionar este problema.
