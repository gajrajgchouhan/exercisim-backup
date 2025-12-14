package lasagna

// define the 'PreparationTime()' function
func PreparationTime(s []string,l int)int{
    if l==0{
         return len(s)*2;
    }
    return len(s)*l;
}

// define the 'Quantities()' function
func Quantities(s []string)(int,float64){
    var n int;
    var se float64;
    for i:=0;i<len(s);i++{
        if(s[i]=="noodles"){
            n++;
        }
		if(s[i]=="sauce"){
            se++;
        }
    }
	return n*50 , se*0.2;
}

// define the 'AddSecretIngredient()' function
func AddSecretIngredient(f []string,o []string){
   o[len(o)-1] = f[len(f)-1]
}

//define the 'ScaleRecipe()' function
func ScaleRecipe(s []float64,n int)[]float64{
	v:=make([]float64,len(s))
    for i:=0;i<len(s);i++ {
        v[i]=s[i]*float64(n)/2;
    }
	return v;
}
