<script setup lang="ts">
import Search from "./Search.vue"
import { ref, Ref } from "vue"
import RootObject from "../services/searchService/interface";

const items = ref()
function selectData(emitData: RootObject) {
    //a la variable reactiva items se le asigna emitData el cual es el body de la peticion que se hizo. la informacion que se necesita de ese body esta dentro de un objeto hits que contiene un atributo tambien llamado hits el cual es un array. los elementos de ese array llamado hits son los que se necesitan mostrar en pantalla en una tabla
    items.value = emitData.hits.hits
}

const body: Ref<string | null> = ref("")
function reqBody(ev:MouseEvent) {
    //al dar click en una fila de la tabla se obtiene el texto de la columna 4 correspondiente a esa fila. dicho texto corresponde a la variable {{ v._source.document.Body }}
    body.value = (<HTMLTableRowElement>ev.currentTarget).childNodes[4].textContent
}
</script>

<template>
    <div>
    <Search class="searchPosition" @reqData="selectData" />

    <table class="styled-table">
        <thead>
            <tr>
                <th>Subject</th>
                <th>From</th>
                <th>To</th>
                <th>Date</th>
            </tr>
        </thead>
        <tbody>
            <!--items contiene un array con los datos que se necesitan mostrar en pantalla en una tabla. como el dato {{ v._source.document.Body }} esta dentro de una etiqueta <template>, entonces dicho dato no se renderiza en la tabla. esto se hace asi porque este dato se necesita almacenarlo en una variable pero no se necesita mostrarlo en la tabla -->
            <tr v-for="v in items" @click="reqBody">
                <td> {{ v._source.document.Subject }}</td>
                <td> {{ v._source.document.From }}</td>
                <td> {{ v._source.document.To }}</td>
                <td> {{ v._source.document.Date }}</td>
                <template>
                    <td> {{ v._source.document.Body }}</td>
                </template>
            </tr>
        </tbody>
    </table>
    <p v-text="body"></p>
</div>
</template>

<style scoped>
.searchPosition{
    position:relative;
    top: 20px;

}
div{
    background-color: #111827;
    height: 1500px;
}
.styled-table {
    border-collapse: collapse;
    margin: 25px 0;
    font-size: 0.9em;
    font-family: sans-serif;
    min-width: 400px;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
}

.styled-table thead tr {
    background-color: #6366f1;
    color: #ffffff;
    text-align: left;
}

.styled-table th,
.styled-table td {
    padding: 12px 15px;
    color: #ffffffee;
}

.styled-table tbody tr {
    border-bottom: 1px solid #6366f1;
    background-color: #1f2937;
}

.styled-table tbody tr:nth-of-type(even) {
    background-color: #374151;
}

.styled-table tbody tr:last-of-type {
    border-bottom: 2px solid #6366f1;
}

.styled-table tbody tr:hover {
    /*cuando se haga hover al contenedor div el estilo se aplica a su elemento hijo .circulo*/
    background-color: rgba(99 102 241 / 0.7);
    cursor: pointer;
}

p{
    color: #ffffffee;
}
</style>