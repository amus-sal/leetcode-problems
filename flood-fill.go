
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc] == newColor {
        return image
    }
    return fillRec(image, sr,sc,image[sr][sc],newColor)
}


func fillRec (image [][]int, sr int, sc int,baseColor int,color int)[][]int{
    if sc < 0 || sr < 0 {
        return image
    }
    if sr >= len(image) || sc >= len(image[sr]){
        return image
    }
    if  baseColor != 0 && image[sr][sc] < 0{
         return image 
    }
    if image[sr][sc] == baseColor {
        image[sr][sc] = color
        image = fillRec(image, sr , sc-1,baseColor,color)
        image = fillRec(image, sr , sc+1,baseColor,color)
        image = fillRec(image, sr -1 , sc,baseColor,color)
        image = fillRec(image, sr  +1, sc,baseColor,color)
        return image
    }
    return image
}