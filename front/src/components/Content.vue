<script setup lang="ts">
import Search from "./Search.vue"
import { ref, Ref } from "vue"

const items = ref()
function selectData(reqData) {
    items.value = reqData.hits.hits
}

const body: Ref<string> = ref("")
function reqBody(ev) {
    body.value = ev.currentTarget.childNodes[4].textContent
}
</script>

<template>
    <div>
    <Search @reqData="selectData" />

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
            <tr v-for="v in items" @click="reqBody">
                <td> {{ v._source.Subject }}</td>
                <td> {{ v._source.From }}</td>
                <td> {{ v._source.To }}</td>
                <td> {{ v._source.Date }}</td>
                <template>
                    <td> {{ v._source.Body }}</td>
                </template>
            </tr>
        </tbody>
    </table>
    <p v-text="body"></p>
</div>
</template>

<style scoped>

Search{
    
}

div{
    background-color: #111827;
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