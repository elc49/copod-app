package com.lomolo.copodapp.ui.common

import androidx.compose.foundation.layout.size
import androidx.compose.material3.Badge
import androidx.compose.material3.BadgedBox
import androidx.compose.material3.Icon
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import androidx.navigation.NavDestination.Companion.hierarchy
import com.lomolo.copodapp.R

sealed class Screen(
    val name: Int,
    val defaultIcon: Int,
    val activeIcon: Int,
    val route: String,
    var showBadge: Boolean = false,
) {
    data object Explore : Screen(
        R.string.explore,
        R.drawable.explore_outlined,
        R.drawable.explore_filled,
        "explore",
        false,
    )
    data object Land: Screen(
        R.string.land,
        R.drawable.global_africa_outlined,
        R.drawable.global_africa_filled,
        "land",
        false,
    )
}


@Composable
fun BottomNavBar(
    currentDestination: NavDestination,
    onNavigateTo: (String) -> Unit,
) {
    val navItems = listOf(Screen.Explore, Screen.Land)

    NavigationBar {
        navItems.forEachIndexed { _, item ->
            // TODO read from current destination
            val isActive = currentDestination.hierarchy.any { it.route == item.route } == true

            NavigationBarItem(selected = false,
                onClick = { if (!isActive) onNavigateTo(item.route) },
                icon = {
                    if (item.showBadge) {
                        BadgedBox(badge = { Badge() }) {
                            Icon(
                                painterResource(if (isActive) item.activeIcon else item.defaultIcon),
                                modifier = Modifier.size(28.dp),
                                contentDescription = null,
                            )
                        }
                    } else {
                        Icon(
                            painterResource(if (isActive) item.activeIcon else item.defaultIcon),
                            modifier = Modifier.size(28.dp),
                            contentDescription = null,
                        )
                    }
                },
                label = {
                    Text(
                        stringResource(item.name),
                        fontWeight = if (isActive) FontWeight.ExtraBold else FontWeight.Normal,
                    )
                })
        }
    }
}