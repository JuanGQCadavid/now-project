import UIKit
import Flutter
import GoogleMaps
import Darwin

@UIApplicationMain
@objc class AppDelegate: FlutterAppDelegate {
  override func application(
    _ application: UIApplication,
    didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?
  ) -> Bool {
    let noApiKey = "NO_API_KEY_SET_UP"
    let apiKey = ProcessInfo.processInfo.environment["GOOGLE_MAP_API_KEY"] ?? noApiKey 
    print(apiKey)
    if apiKey == noApiKey {
      print("**************")
      print("You sould create a Variable envairoment with the name GOOGLE_MAP_API_KEY pointing to the google api key")
      print("**************")
    }
    
    GMSServices.provideAPIKey("")
    GeneratedPluginRegistrant.register(with: self)
    return super.application(application, didFinishLaunchingWithOptions: launchOptions)
  }
}
