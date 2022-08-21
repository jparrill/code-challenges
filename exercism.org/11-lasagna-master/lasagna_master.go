package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(xl []string, time int) int {
    if time <= 0 {
        time = 2
    }

    return len(xl) * time
}

// TODO: define the 'Quantities()' function
func Quantities(xi []string) (int, float64) {
    var qNoodels int
    var qSauce float64

    for _,v := range xi {
        if v == "noodles" {
            qNoodels += 50
        } else if v == "sauce" {
        	qSauce += 0.2
        }
    }

    return qNoodels, qSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    si := friendsList[len(friendsList) - 1]
    myList[len(myList) - 1] = si
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(xf []float64, portions int) []float64 {
    sxf := make([]float64, len(xf))
    copy(sxf, xf)

    for i,v := range xf {
        sxf[i] = v * (float64(portions)/2)
    }

    return sxf
}
