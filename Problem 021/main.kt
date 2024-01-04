
fun main(args : Array<String>) {

    // Init
    val max : Int = 10000;

    // Find Amicables
    var amicables : IntArray = IntArray(0)

    for (i in 0 until max-1){ // i to N-1
        for (j in i+1 until max){ // i+1 to N
            if (i != j) { // exclude i = j
                if (spd(fpd(i)) == j && spd(fpd(j)) == i) { // where all magic happens
                    if (!amicables.contains(i)) amicables += i
                    if (!amicables.contains(j)) amicables += j
                }
            }
        }
    }

    // Sum Amicables
    var sum : Int = 0;
    for (x in amicables){
        sum += x
    }

    println(sum)
}

fun fpd(n : Int) : IntArray { // find_proper_divisors
    var divs : IntArray = IntArray(0)
    for (i in 1 until n){
        if (n % i == 0){
            divs += i
        }
    }
    return divs
}

fun spd(arr : IntArray) : Int{ // sum_proper_divisors
    var sum : Int = 0;
    for (i in arr){
        sum += i
    }
    return sum
}