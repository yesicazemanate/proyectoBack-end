"use client"
import React, { useState,useEffect } from 'react'
import axios from 'axios'
import Card from '@/components/Card'
import Image from 'next/image'
import '@/app/style.css'


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
   <header>
        <p className="titulo titulos">RESTAURANTE</p>
        </header>
        <div className="contenedorPrincipal"> 
        
        <div className="contenido1">
        <form className="formu">
            <fieldset>
              <h1>Ingrese Informacion de alimentos</h1>
                <legend className='contenedorTitulo'></legend>
                <input required type="text" placeholder="Ingrese Nombre" />
                <input required type="text" placeholder="Ingrese Descripcion" />
                <input required type="text" placeholder="Ingrese Imagen" />
                <button type="submit" className="boton" >Enviar</button>
               </fieldset>
              
        </form>
        </div>
        <div className="contenido">
        <p className="historial titulos">Historial de Alimentos</p>
        <div className="linea"></div>
        <div className="container">
       { alimentoItems.map(data=>
        <div key={data._id} className='contenedorProducto'>
         <h1 className='name'> {data.name}</h1> 
       <Image src={data.image} width={150} height={150}  loader={({ src }) => src} alt='imagen de sancocho' className='image'/>
        <p className='descripcion'>{data.descripcion}</p>
        </div>
      )}
        </div>
        </div>
        </div>

        </>
    )
}

    

 
