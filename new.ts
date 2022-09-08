for (i=0;i<s.length;i++){
  if (s[i].match(/[a-zA-Z]/i)){
     let code=s[i].charCodeAt(0)
      if (code>=65 && code<=90){
          //upercase
          if (k>26){
              k = k % 26        
          }    
          let newcode = code + (code + k>90 ? k-26:k); 
          ret= ret.concat(String.fromCodePoint(newcode))
          console.log(s[i],String.fromCodePoint(newcode))
      }else if (code>=97 && code<=122){
          //lowercase
          if (k>26){
              k = k % 26        
          }    
              let newcode = code +(code + k>122? k-26:k);
              ret= ret.concat(String.fromCodePoint(newcode))
              console.log(s[i],String.fromCodePoint(newcode))
      }
  }else {
      ret=ret+s[i]
      console.log(s[i])
  }
}

//check if char 