<template>
  <div class="">
    
  <v-app id="inspire">
    
    <v-row style="flex:none">
      <v-col cols="2">
       <v-menu ref="menu1" v-model="menu1" :close-on-content-click="false" :return-value.sync="date1" transition="scale-transition" offset-y min-width="auto">
          <template v-slot:activator="{ on, attrs }">
            <v-text-field v-model="date1" label="Start Date" prepend-icon="mdi-calendar" readonly v-bind="attrs" v-on="on"></v-text-field>
          </template>
          <v-date-picker v-model="date1" no-title scrollable>
            <v-spacer></v-spacer>
            <v-btn text color="primary" @click="menu1 = false">
              Cancel
            </v-btn>
            <v-btn text color="primary" @click="$refs.menu1.save(date1)">
              OK
            </v-btn>
          </v-date-picker>
        </v-menu>
      </v-col>
      <v-spacer></v-spacer>
      <v-col cols="2">
        <v-menu ref="menu" v-model="menu" :close-on-content-click="false" :return-value.sync="date" transition="scale-transition" offset-y min-width="auto">
          <template v-slot:activator="{ on, attrs }">
            <v-text-field v-model="date" label="End Date" prepend-icon="mdi-calendar" readonly v-bind="attrs" v-on="on"></v-text-field>
          </template>
          <v-date-picker v-model="date" no-title scrollable>
            <v-spacer></v-spacer>
            <v-btn text color="primary" @click="menu = false">
              Cancel
            </v-btn>
            <v-btn text color="primary" @click="$refs.menu.save(date)">
              OK
            </v-btn>
          </v-date-picker>
        </v-menu>
      </v-col>
      <v-col cols="1" style="margin:auto 0">
        <v-btn depressed color="primary" v-on:click="postreq()">
          Search
        </v-btn>
      </v-col>
      <v-col cols="7"></v-col>
    </v-row>

    <v-card>
      <v-card-title>
        Orders
        <v-spacer></v-spacer>
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="OrderInfo"
        :items-per-page="5"
        :search="search"
      ></v-data-table>
    </v-card>

  </v-app>
  </div>
</template>

<script>
// //Bootstrap and jQuery libraries
// import 'bootstrap/dist/css/bootstrap.min.css';
// import 'jquery/dist/jquery.min.js';
// //Datatable Modules
// import "datatables.net-dt/js/dataTables.dataTables"
// import "datatables.net-dt/css/jquery.dataTables.min.css"

//import $ from 'jquery'; 
import axios from 'axios';

export default {
 name: 'Orders',
  mounted(){
    //API Call
    axios.get("http://localhost:8090/orders").then((res)=>
    {
      this.OrderInfo = res.data;
      //$('#example').DataTable();
    })
  },
  data: function() {
        return {
            // dates: ['2019-09-10', '2019-09-20'],
            search: '',
            headers: [
              {
                text: 'Order Name',
                align: 'start',
                filterable: true,
                value: 'order_name',
              },
              { text: 'Product', value: 'product', align: ' d-none' }, // ' d-none' hides the column but keeps the search ability
              { text: 'Company Name', value: 'company_name' },
              { text: 'Customer Name', value: 'name' },
              { text: 'Created At', value: 'created_at' },
              { text: 'Delivered Quantity', value: 'delivered_quantity' },
            ],
            OrderInfo:[]
        }
    },
    methods: {
      postreq: function() {
        var data = {"StartDate": this.date1, "EndDate": this.date}
        /*eslint-disable*/
        console.log(data) 
        /*eslint-enable*/
        axios({ method: "POST", url: "http://localhost:8090/ordersSearch", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
            // this.response = result.data;
            /*eslint-disable*/
            this.OrderInfo = result.data;
            console.log(result.data) 
          /*eslint-enable*/
          }).catch( error => {
            /*eslint-disable*/
              console.error(error);
            /*eslint-enable*/
        });
      }
    }
}
</script>