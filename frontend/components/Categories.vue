<template>
  <div>
    <div class="container my-12 mx-auto px-4 md:px-12 max-w-md">
      <div class="flex flex-wrap lg:-mx-16">
        <div
          class="my-5 px-4 w-full md:w-1/2 cursor-pointer"
          v-for="(category,idx) in Categories" :key="idx">
          <article class="overflow-hidden rounded-lg" @click.prevent="openCatalog(category)">
            <img :src = "`http://localhost:8080/static/` + category.image_url" alt="Placeholder" class="block  w-full" />
            <header class="flex items-center justify-between leading-tight p-4 md:p-4 bg-gray-100">
              <h1 class="text-lg"><p class="text-black font-medium md:items-center">{{category.name}}</p></h1>
            </header>
          </article>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import gql from 'graphql-tag';
const Categories_QUERY = gql`
    query {
      Categories{
        id
        name
        image_url
      }
    }
`
export default {
    layout: 'CatalogLayout',
    apollo: {
        Categories: {
            query: Categories_QUERY,
        }, 
    },
    methods:{
        openCatalog(category){
          this.$router.push('/catalog/'+ category.id)
        },
    },
}
</script>
