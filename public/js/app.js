var demoApp = angular.module("demoApp", [])

demoApp.controller("demoController", function($scope, $http) {
  $scope.items = [];

  $scope.loadItems = function() {
    $http.get('/api/demo').success(function(data) {
      $scope.items = data
    })
  }

  $scope.loadItem = function(itemPath) {
    $http.get(itemPath).success(function(data) {
      $scope.items = []
      $scope.item = data
    })
  }

})
