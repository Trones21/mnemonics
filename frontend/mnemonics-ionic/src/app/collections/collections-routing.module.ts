import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CollectionsPage } from './collections.page';
import { newCollectionPage } from './newCollection.page';
import { SingleCollectionPage } from './singleCollection.page';

const routes: Routes = [
  {
    path: '',
    component: CollectionsPage,
  },
  {
    path:'new',
    component: newCollectionPage
  },
  {
    path:'details/:id',
    component: SingleCollectionPage
  }
  ];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CollectionsPageRoutingModule {}