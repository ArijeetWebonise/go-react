import * as firebase from "firebase";
import axios from 'axios';

const InitNotification = {
  notification: {
    title: "rishithegamecreator",
    body: "These is body",
    click_action: "https://rishithegamecreator-1054.firebaseapp.com",
  },
  to: ""
};

const SERVER_KEY = 'AAAAGdnhHIY:APA91bGd6cx-TPxalfo8MsNhiv3JBdbPf8RwCduBc-0FAIeZxeCKxKP3x2JCOs7_F89fzkImZ11DndTA78ySprUg-6sfPxXpG9fR64pzbLnaViZCfNvmODQF2fRxJTJTnaH6ARPmcYHJ';

const keys = {
  apiKey: "AIzaSyDB_B-chtnaDDlvThaGFYRw2MhEe3N-AAc",
  authDomain: "rishithegamecreator-1054.firebaseapp.com",
  databaseURL: "https://rishithegamecreator-1054.firebaseio.com/",
  projectId: "rishithegamecreator-1054",
  storageBucket: "rishithegamecreator-1054.appspot.com",
  messagingSenderId: "111029591174"
};

firebase.initializeApp(keys);

export default class Firebase {
  static LoginViaGoogle() {
    const googleAuthProvider = new firebase.auth.GoogleAuthProvider();
    return firebase.auth().signInWithPopup(googleAuthProvider);
  }

  static LoginViaForm(email, pass) {
    return firebase.auth().signInWithEmailAndPassword(email, pass);
  }

  static LoginViaGithub() {
    const githubAuthProvider = new firebase.auth.GithubAuthProvider();
    return firebase.auth().signInWithPopup(githubAuthProvider);
  }

  static UserPasswordReset(email) {
    return firebase.auth().sendPasswordResetEmail(email);
  }

  static getUserDetail(userID) {
    const userRef = firebase.database().ref().child('user/');
    return userRef.orderByChild('uid').equalTo(userID).once('value');
  }

  static RegisterUser(email, password, firstname, lastname) {
    return new Promise((resolve, reject) => {
      firebase.auth().createUserWithEmailAndPassword(email, password).then((user) => {
        user.user.sendEmailVerification()
          .then(() => {
            console.log("sending mail verification");
          });
        const userRef = firebase.database().ref().child('user/');
        userRef.push({
          uid: user.user.uid,
          firstname,
          lastname,
        })
        .then((snapshot) => {
          resolve(snapshot);
        })
        .catch((error) => {
          reject(error);
        });
      })
      .catch((error) => {
        reject(error);
      });
    });
  }

  static onAuthStateChanged(callback) {
    return firebase.auth().onAuthStateChanged(user => callback(user));
  }

  static Logout() {
    return firebase.auth().signOut()
  }

  static UserserviceWorker(registration) {
    firebase.messaging().useServiceWorker(registration);
  }

  static askForPermissionToReceiveNotifications() {
    return new Promise((resolve, reject) => {
      const requestPermission = async () => {
        try {
          const messaging = firebase.messaging();
          await messaging.requestPermission();
          const token = await messaging.getToken();
          console.log('notification token:', token);
          resolve(token);
        } catch(error) {
          reject(error);
        }
      };

      requestPermission();
    });
  }

  static sendMessage(data = InitNotification) {
    axios.post('https://fcm.googleapis.com/fcm/send', data, { headers: [
      { Authorization: `key: ${SERVER_KEY}`}
    ]});
  }
}
