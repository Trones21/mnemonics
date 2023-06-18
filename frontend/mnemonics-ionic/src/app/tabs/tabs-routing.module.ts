import { NgModule } from '@angular/core';
import { RouterModule, Routes, ActivatedRoute, OnSameUrlNavigation, ExtraOptions } from '@angular/router';
import { TabsPage } from './tabs.page';
import { CategoryPage } from '../category/category.page';
import { CollectionCard } from '../components/collectionCard';
import { MnemonicCard } from '../components/mnemonicCard';
import { SingleCollectionPage } from '../collections/singleCollection.page';
import { ProfilePage } from '../profile/profile.page';
import { EditorPage } from '../editor/editor.page';
import { ProfileService } from '../profile/profile.service';
import { DataService } from '../data.service';
import { PracticeSetupPage } from '../practice/practiceSetup.page';
import { Sandbox } from '../helpers/sandbox.page';



const routes: Routes = [
  {
    path: '',
    component: TabsPage,
    children: [
      {
        path: 'practice',
        loadComponent:() => import('../practice/practiceSetup.page').then(m => m.PracticeSetupPage)
      },
      {
        path: 'editor',
        loadComponent: () => import('../editor/editor.page').then(m => m.EditorPage)
      },
      
      {
        path: 'collections',
        loadChildren: () => import('../collections/collections.module').then(m => m.CollectionsPageModule)
      },
      {
        path: 'profile/:id',
        loadComponent: () => import('../profile/profile.page').then(m => m.ProfilePage)
      },

      //Place components here to see them in isolation
      {
        path: 'test',
        component: Sandbox
      },
      {
        path: '',
        redirectTo: '/collections',  //Basically which tab show be selected when first loading / empty path
        pathMatch: 'full'
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
})
export class TabsPageRoutingModule {}
