<template>
      <div class="center">
      <va-card color="white" gradient style="max-width: 600px; min-width: 400px">
        <div class="badges">
        <va-badge color="#E7FFDC" :text="petition.tags[0]" class="mr-4"> </va-badge>
        <va-badge color="#E7FFDC" :text="petition.tags[1]" class="mr-4"> </va-badge>
        <va-badge color="#E7FFDC" :text="petition.tags[2]" class="mr-4"> </va-badge>
        </div>
        <router-link
            :to="{ name: 'petition', params: { id: petition.id } }" style="text-decoration: none; color: black;">
        <h6 class="text">{{petition.title}}</h6>
        </router-link>
        <va-progress-bar
            class="mt-2"
            :model-value="value"
            size="1.5rem"
            :color="color"
            content-inside
            show-percent
        />
        <div class="additional-info">
          <span class="mr-4">{{ petition.authorName }}</span>
          <span class="mr-4">{{ petition.dateOfExpiration }}</span>
          <span class="mr-4">{{petition.numberOfSignatures}}/{{petition.signaturesTarget}}</span>
        </div>
      </va-card>
    </div>
</template>
<script>

export default {
  name: 'PetitionCard',
  props: {
    petition: {
      id: Number,
      title: String,
      authorName: String,
      authorEmail: String,
      dateOfCreation: String,
      dateOfExpiration: String,
      tags: Array,
      petitionContent: String,
      numberOfSignatures: Number,
      signaturesTarget: Number
    }
  },
  computed: {
    value() {
      return this.petition.numberOfSignatures / this.petition.signaturesTarget * 100
    },
    color(){
      if (this.value < 50) {
        return 'danger'
      } else if (this.value < 100) {
        return 'warning'
      } else {
        return 'success'
      }
    },
  }
}
</script>

<style>
.additional-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-left: 1%;
}

.text
{
  margin-left: 1%;
  margin-bottom: 1.5%;
  font-size: medium;
  font-weight: bold;
}

.badges {
  margin-left: 1%;
  margin-bottom: 1.5%;
}


.center {
  margin-left: auto;
  width: 90%;
  padding: 12px;
}
</style>