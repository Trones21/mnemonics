import { Component } from '@angular/core';
import { CategoriesService } from '../category/categories.service';
import { CategoryDetails } from '../category/category.page';
import { CollectionInfo } from '../components/collectionCard';


@Component({
  selector: 'app-collections',
  templateUrl: 'collections.page.html',
})
export class CollectionsPage {
  loggedInUserID = localStorage.getItem('loggedInUserId')
  view: string = 'Search'
  
  categories: Array<Category> = []
  selectedCategoryView: boolean = false;
  selectedCategory: CategoryDetails | null = null
  constructor(private categoriesService: CategoriesService) {

  }

  showSearch(){
    this.view = 'Search'

  }

  showBrowse(){
    this.view = 'Browse'
  //   this.categories = [{
  //     id: '1',
  //     name: 'Medical',
  //     collectionCount: 14 
  // },
  // {
  //     id: '2',
  //     name: 'Math',
  //     collectionCount: 2 
  // },
  // {
  //     id: '3',
  //     name: 'Biology',
  //     collectionCount: 4 
  // },
  // ]
    this.categoriesService.getCategories().then(c => this.categories = c)
  }

  Search(event: any){
    console.error('Need to Implement')
    console.log('search', event.target)

  }

  selectCategory(categoryId: string){
    this.categoriesService.getCategoryByID(categoryId).then(c => {
      console.log('after get', c)
      //Rename Property to match collectionInfo
      c.collectionCards = c.collections;
      delete c.collections;
      this.selectedCategory = (c as CategoryDetails)
      this.selectedCategoryView = true;
    })
  }

}

interface Category{
  id: string,
  name: string,
  collectionCount: number
}