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
    GMSServices.provideAPIKey("AIzaSyDqVfIg2SaH7I2DONpg-t5wY1EwjDXq2Vg")
    GeneratedPluginRegistrant.register(with: self)
    return super.application(application, didFinishLaunchingWithOptions: launchOptions)
  }
}
