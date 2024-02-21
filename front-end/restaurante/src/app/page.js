"use client"
import React from 'react'
import { useEffect, useState } from 'react'

const restaurant=()=> {
 const [alimentoItems, setAlimentoItems]= useState([])
useEffect(() =>{
  const consultarApi = async() =>{
    try{
    const url = 'http://localhost:8080/alimentos'
    const respuesta = await fetch(url)
    const resultado =await respuesta.json()
    setMenuItems(resultado);
  } catch (error) {
    console.error('Error al obtener los alimentos:', error);
  }
  }
  consultarApi()
},[])

  return (
    <>
<ul>
  
    {alimentoItems.map(item=>{
      <li key={item.id}> {item.name} {item.descripcion}

      </li>
    })

    }

    
 
</ul>
    </>
  )
}
export default restaurant

