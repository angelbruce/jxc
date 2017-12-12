import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { Geofence } from '@ionic-native/geofence';

@Component({
  selector: 'page-about',
  templateUrl: 'about.html'
})
export class AboutPage {
  latitude
  longitude
  radius
  constructor(public navCtrl: NavController, private geofence: Geofence) {
    geofence.initialize().then(() => {
      console.log('Geofence Plugin Ready');
    }, (err) => console.log(err));
  }


  private addGeofence() {
    //options describing geofence
    let fence = {
      id: '69ca1b88-6fbe-4e80-a4d4-ffad374aacda', //any unique ID
      latitude:       37.285951, //center of geofence radius
      longitude:      -121.936650,
      radius:         100, //radius to edge of geofence in meters
      transitionType: 3, //see 'Transition Types' below
      notification: { //notification settings
          id:             1, //any unique ID
          title:          'You crossed a fence', //notification title
          text:           'You just arrived to Gliwice city center.', //notification body
          openAppOnClick: true //open app when notification is tapped
      }
    }
  
    this.geofence.addOrUpdate(fence).then(
       () => console.log('Geofence added'),
       (err) => console.log('Geofence failed to add')
     );
    }
}
