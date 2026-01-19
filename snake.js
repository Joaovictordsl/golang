function escolheTaxi(tf1,vqr1,tf2,vqr2) {
  
  tf1 = Number(tf1);
  vqr1 = Number(vqr1);
  tf2 = Number(tf2);
  vqr2 = Number(vqr2);
  
//   if (tf1 < 0 || vqr1 < 0 || tf2 < 0 || vqr2 < 0) {
//     throw new Error("negative numbers.");
//   }
  
  if(tf1 === tf2 && vqr1 === vqr2) return "Tanto faz";
  
  let d;
  //let taxi1 = tf1 + (vqr1 * d); 
  //let taxi2 = tf2 + (vqr2 * d);
  
  //taxi1 = tf1 + (vqr1 * d);
  //taxi1 - tf1 = vqr1 * d
  //d = (taxi1 - tf1)/vqr1 
  
  
  //taxi2 = tf1 + (vqr2 * d);
  //taxi2 - tf1 = vqr2 * d
  //d = (taxi2 - tf2)/vqr2 
  
  d = (tf2 - tf1) / (vqr1 - vqr2);
  
  if (d <= 0){
    return vqr1 < vqr2 ? "Empresa 1" : "Empresa 2";
  }
  
  let taxi1, taxi2;
  
  if (vqr1 > vqr2){
    taxi1 = "1";
    taxi2 = "2";
  }
  
  return `Empresa 1 quando a distância < ${d}, Tanto faz quando a distância = ${d}, Empresa 2 quando a distância > ${d}`;
  
}

console.log(escolheTaxi("2.5","1.0","5.0","0.75"))