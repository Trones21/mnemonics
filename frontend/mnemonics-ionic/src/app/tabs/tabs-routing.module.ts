import { NgModule } from '@angular/core';
import { RouterModule, Routes, ActivatedRoute, OnSameUrlNavigation, ExtraOptions } from '@angular/router';
import { TabsPage } from './tabs.page';
import { CategoryPage } from '../category/category.page';
import { CollectionCard } from '../components/collectionCard';
import { MnemonicCard } from '../components/mnemonicCard';
import { SingleCollectionPage } from '../collections/singleCollection.page';
import { ProfilePage } from '../profile/profile.page';
import { ProfileService } from '../profile/profile.service';
import { DataService } from '../data.service';
import { Sandbox } from '../helpers/sandbox.page';



const routes: Routes = [
  {
    path: '',
    component: TabsPage,
    children: [
      {
        path: 'swiper',
        loadChildren: () => import('../swiper/swiper.module').then(m => m.SwiperPageModule)
      },
      {
        path: 'editor',
        loadChildren: () => import('../editor/editor.module').then(m => m.EditorPageModule)
      },
      
      {
        path: 'collections',
        loadChildren: () => import('../collections/collections.module').then(m => m.CollectionsPageModule)
      },
      {
        path: 'profile',
        loadChildren: () => import('../profile/profile.module').then(m => m.ProfilePageModule)
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
