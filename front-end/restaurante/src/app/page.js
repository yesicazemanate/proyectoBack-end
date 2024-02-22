"use client"
import React, { useState,useEffect } from 'react'
import axios from 'axios'
import Card from '@/components/Card'



export default function Home() {
  const [alimentoItems, setAlimentoItems]=useState([])
useEffect(()=>{
  const conectarApi= async()=>{
    try{
      const response = await axios.get('http://localhost:8080/alimentos')
      const data = await response.data
 //console.log(data)
      setAlimentoItems(data)
    }catch(error){
     console.error(error)
  
    }}
   conectarApi()
  })
    return (
      <>
      {alimentoItems.forEach(data=>{
        <div key={data._id}> name:{data.name}</div>
      })}
      <ul> hola

 </ul>
  </>
    )
  }
  
    

 
