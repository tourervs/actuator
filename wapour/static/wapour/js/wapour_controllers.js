wapourApp.controller('appController',['settingsService','dashboardDataService','websocketService','$scope', function initController(settingsService,dashboardDataService,websocketService, $scope) {

    $scope.navigation_menu = {} ;
    $scope.initApp = function(settings) {
        settingsService.setSettings(settings);
        app_settings = settingsService.getSettings();
        if (app_settings.websocket == true){
            websocketService.createWsConnection(settings["ws_url"]);
        }
        app_data_url = app_settings.app_data_url ; 


        var data_promise = dashboardDataService.GetHttp(app_data_url);
        data_promise.then(function(result) {
            if (result.status == "ok") {
                $scope.navigation_menu = result.data.navigation_menu; 
            }
        });


        //console.log(app_data);
        //if (app_data.status == "ok") {
        //    $scope.navigation_menu = app_data.navigation_menu ; 
        //}
    };
    $scope.$on("$destroy", function(){
        
        websocketService.closeWsConnection();
        console.log("Exit from initController . Closing ws-connetion");
    });
    /*$scope.$on('$routeChangeStart', function(){
        websocketService.closeWsConnection();
        alert("Catching routeChangeStart");
        console.log("Exit from initController  via routeChangeStart . Closing ws-connetion");
    });*/
    
}]);

wapourApp.controller('mainController', ['$scope','websocketService', function($scope, websocketService) {
    var ws_connect_retries = 10 ; 
    var data            = {};
    var message         = {"datatype":"message_chat"};
    
    data['author']      = "user1";
    data['message']     = "Hello all!";
    message["data"]     = data
    
    var send_test_message  = function(){websocketService.sendRequest(message)};
    var data               = websocketService.wsReady(ws_connect_retries, send_test_message);

    $scope.dashboard_select = function(dashboard_name) {
        alert("Selected dashboard:"+dashboard_name); 
    }
    $scope.dashboards_list = function() {

    }
}]);

wapourApp.controller('dashboardController', ['$scope','dashboardDataService', function($scope, dashboardDataService) {

    $scope.dashboard_data = {} ;

    var notifier = function(dashboard_id, dashboard_group_id) {
        console.log("::calling callback::");
        //$scope.dashboard_data = dashboardDataService.GetDashboardData(dashboard_id, dashboard_group_id); 
        var data_promise = dashboardDataService.GetDashboardData(dashboard_id, dashboard_group_id);
        data_promise.then(function(result) {
           if (result.status == "ok") {
               console.log("Result data");
               console.log(result.data) ; 
           }
           return result;
        });
        console.log("-- dashboard_data --")
        console.log($scope.dashboard_data);
    }
    var notifier_url = function(dashboard_url) {
        $scope.dashboard_data = dashboardDataService.GetDashboardDataByUrl(dashboard_url);
        console.log($scope.dashboard_data);
    }
    dashboardDataService.AddCallback(notifier);



}]);

