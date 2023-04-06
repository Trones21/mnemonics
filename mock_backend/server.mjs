//Common JS import * from directory? 
//I have a little mock backend, each entity has it's own file in a data folder 
//is it possible to do something like this?
//import * as entities from './data/*.mjs'
//Note: I suppose I could put all entities in one file, then export {...} / import *, but I prefer to keep them in separate files  

import mnemonic from  './data/mnemonic.mjs'
import profile from  './data/profile.mjs'
import collection from  './data/collection.mjs'
import category from  './data/category.mjs'
import categories from  './data/categories.mjs'
import express from 'express'
import cors from 'cors'
const app = express()
const port = 4000
const angularPort = 8100

app.use(express.json()) 

app.options('/*', cors())

//Working
//Note: This is a catch-all, i will have different endpoints for different entities
app.post('/*', cors(), function(req, res){
  console.log(req.body)
  //let x = JSON.parse(req.body)
  //console.log(x)
  res.json({"ReceivedObject":true})
})

app.get('/profile/:id', function (req, res) {
  console.log('get profile')
  res.set('Access-Control-Allow-Origin', `http://localhost:${angularPort}`);
  res.json(profile);
})

app.get('/categories/', cors(), function (req, res) {
  res.json(categories);
})
app.get('/category/:id', cors(), function (req, res) {
  console.log('send res', category)
  res.json(category);

})


//To Test
app.get('/mnemonics?filter=:filterLogic', cors(), function (req, res) {
  
  //res.json(mnemonics);
})

app.get('/collection/:id', function (req, res) {
  res.json(collection);
})

app.get('/collections?categoryId=:id', cors(), function (req, res) {
  //Planned architecture categoryId is an attribute of the collection -- you cannot put a collection in multiple categories
  
  //res.json(collections);
})



app.listen(port, () => console.log(`Example app listening on port ${port}!`))