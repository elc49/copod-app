package com.lomolo.copodv2.ui.screens

import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Badge
import androidx.compose.material3.BadgedBox
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.lomolo.copodv2.R
import com.lomolo.copodv2.ui.navigation.Navigation
import com.lomolo.copodv2.viewmodels.MainViewModel

object DashboardScreenDestination : Navigation {
    override val title = null
    override val route = "dashboard"
}

sealed class Screen(
    val name: Int,
    val defaultIcon: Int,
    val activeIcon: Int,
    val route: String,
    var showBadge: Boolean = false,
    val childRoute: List<String> = listOf(),
) {
    data object Explore : Screen(
        R.string.explore,
        R.drawable.explore_outlined,
        R.drawable.explore_filled,
        "explore",
        false,
    )
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun DashboardScreen(
    modifier: Modifier = Modifier,
    mainViewModel: MainViewModel,
    onNavigateTo: (String) -> Unit,
) {
    var openDialog by remember { mutableStateOf(false) }
    val navItems = listOf(Screen.Explore)

    Scaffold(topBar = {
        TopAppBar(title = {}, actions = {
            Image(
                painter = painterResource(R.drawable._9872287),
                contentScale = ContentScale.Crop,
                modifier = Modifier
                    .size(48.dp)
                    .clip(CircleShape)
                    .clickable { openDialog = true },
                contentDescription = stringResource(R.string.user),
            )
        })
    }, bottomBar = {
        NavigationBar {
            navItems.forEachIndexed { _, item ->
                // TODO read from current destination
                val isActive = false

                NavigationBarItem(selected = false, onClick = { onNavigateTo(item.route) }, icon = {
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
                }, label = {
                    Text(
                        stringResource(item.name),
                        fontWeight = if (isActive) FontWeight.ExtraBold else FontWeight.Normal,
                    )
                })
            }
        }
    }) { innerPadding ->
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(innerPadding)
        ) {
            if (openDialog) {
                AccountDetails(setDialog = { openDialog = it })
            }
        }
    }
}