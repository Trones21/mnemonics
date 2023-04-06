import { Injectable } from '@angular/core';
import { DataService } from '../data.service';

@Injectable({
  providedIn: 'root'
})
export class CategoriesService {
  baseUrl: string
  constructor(dataService: DataService,) { 
     this.baseUrl = dataService.baseUrl;
  }

  async getCategories(){
    let res = await fetch(`${this.baseUrl}/categories/`)
    let obj = await res.json();
    console.log(obj)
    return obj;
  }

  
  async getCategoryByID(categoryId: string){
    let res = await fetch(`${this.baseUrl}/category/${categoryId}`)
    let obj = await res.json();
    console.log(obj)
    return obj.category;
  }
}