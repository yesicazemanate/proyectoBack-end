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
    console.log('hola si sirvio')
   // setAlimentoItems(resultado)
    }catch(error){
      console.error('error al obtener los alimentos', error)
    }
  }
  consultarApi()
},[])

  return (
    <>

    </>
  )
}
export default restaurant

