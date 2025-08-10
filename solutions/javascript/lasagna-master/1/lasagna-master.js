/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */
let noodlesLayer = 50
let souceLayer = 0.2


export function cookingStatus(timer) {
  if (timer === 0){
    return 'Lasagna is done.'
  } else if (timer > 0) {
    return 'Not done, please wait.'
  } else {
    return 'You forgot to set the timer.'
  }
}

export function preparationTime(array, timer =2) {
  return array.length * timer
}

export function quantities(layers) {
  const objResult = {noodles: 0, sauce: 0};

  for(let layer of layers) {
    switch(layer) {
      case 'noodles':
          objResult[layer] += 50
          break
      case 'sauce':
        objResult[layer] += 0.2
        break
    }
  }
  return objResult
}


export function addSecretIngredient (friendsList, myList) {
  myList.push(friendsList[friendsList.length-1])
}

export function scaleRecipe(recipe, portions) {
  const scaled = { ...recipe }
  const factor = portions / 2
  for (let ingredient in scaled) {
    scaled[ingredient] *= factor
  }
  return scaled
}
