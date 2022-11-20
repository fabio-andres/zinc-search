<script setup lang="ts">
import { ref, Ref, onMounted } from "vue"

const inputSearchElem: Ref<null> = ref(null)
onMounted(() => {
    inputSearchElem.value
})

const emit = defineEmits(['reqData'])

interface RootObject {
  took: number;
  timed_out: boolean;
  _shards: Shards;
  hits: Hits;
}

interface Hits {
  total: Total;
  max_score: number;
  hits: Hit[];
}

interface Hit {
  _index: string;
  _type: string;
  _id: string;
  _score: number;
  '@timestamp': string;
  _source: Source;
}

interface Source {
  Body: string;
  'Content-Transfer-Encoding': string;
  'Content-Type': string;
  Date: string;
  From: string;
  'Message-ID': string;
  'Mime-Version': string;
  Subject: string;
  To: string;
  'X-FileName': string;
  'X-Folder': string;
  'X-From': string;
  'X-Origin': string;
  'X-To': string;
  'X-bcc': string;
  'X-cc': string;
}

interface Total {
  value: number;
}

interface Shards {
  total: number;
  successful: number;
  skipped: number;
  failed: number;
}

const req = async () => {
    const user: string = "admin"
    const password: string = "Complexpass#123"
    const encodeCredentials: string = btoa(user + ":" + password)
    const url: string = 'http://localhost:4080/es/enron1/_search'
    const reqOptions = {
        method: "POST",
        headers: {
            "Authorization": "Basic " + encodeCredentials,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "query": {
                "match_phrase": {
                    "Body": inputSearchElem.value.value

                }
            }
        }),
    };

    try {
        const res = await fetch(url, reqOptions)
        const data:Promise<RootObject> = await res.json();
        emit("reqData", data)
    } catch (error) {
        console.log('error al consumir api', error);
    }
}

</script>

<template>
    <section class="main-search">
        <button @click="req"><span class="material-icons-round main-search__search-icon">search</span></button>
        <input class="main-search__input" type="text" ref="inputSearchElem">
    </section>
</template>

<style scoped>
span{
    color: #6366f1;
    font-size: 27px;
}
span:hover {
    box-shadow: 0 2px 12px gray;
    cursor: pointer;
}
.main-search {
    background-color: #111827;
    width: 580px;
    border-radius: 101px;
    border: 1px solid #6366f1;
    margin: 0 auto;
    margin-bottom: 35px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.main-search:hover {
    box-shadow: 0 4px 24px gray;
}

.main-search__input {
    background-color: #111827;
    width: 500px;
    height: 40px;
    border: none;
    outline: none;
    color: #ffffffee;
    font-size: 19px;
}

.main-search__micro-icon {
    background-position: center;
    background-size: contain;
    width: 16px;
    height: 21px;
    cursor: pointer;
    margin: 0;
}

button {
    background-color: #111827;
    border: none;
}

</style>